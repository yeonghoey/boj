package runner

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Runner keeps the context of running a command.
type Runner struct {
	t    *testing.T
	name string
	args []string
}

// New returns an instance of Runner
func New(t *testing.T, name string, arg ...string) *Runner {
	return &Runner{t, name, arg}
}

// Result represents the context of the execution.
type Result struct {
	t      *testing.T
	input  string
	got    string
	err    error
	errmsg string
}

// Run executes "go run ." with input as stdin and
// compares the output with the content of wantOutputName.
func (r *Runner) Run(input string) *Result {
	cmd := exec.Command("go", "run", ".")
	f, err := os.Open(input)
	if err != nil {
		return &Result{r.t, input, "", err, ""}
	}
	defer func() { _ = f.Close() }()
	cmd.Stdin = f

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	errmsg := stderr.String()
	if err != nil {
		return &Result{r.t, input, "", err, errmsg}
	}

	output := string(normalize(stdout.Bytes()))
	got := strings.TrimSpace(output)
	return &Result{r.t, input, got, err, errmsg}
}

// Want tests the result against the content of wantOutputName.
func (r *Result) Want(wantOutputName string) {
	if r.err != nil {
		r.t.Errorf("%q failed: %v\n%s", r.input, r.err, r.errmsg)
		return
	}

	content, err := ioutil.ReadFile(wantOutputName)
	if err != nil {
		r.t.Error(err)
		return
	}

	want := string(normalize(content))
	want = strings.TrimSpace(want)
	if r.got != want {
		r.t.Errorf("%q failed: got %s, want %s", r.input, r.got, want)
		return
	}
}

func normalize(b []byte) []byte {
	// Win -> Unix: replace CR LF with LF
	b = bytes.Replace(b, []byte("\r\n"), []byte("\n"), -1)
	// Mac -> Unix: replace CF with LF
	b = bytes.Replace(b, []byte("\r"), []byte("\n"), -1)
	return b
}
