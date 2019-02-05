package main

import (
	"testing"

	"github.com/yeonghoey/boj/runner"
)

func TestMain(t *testing.T) {
	r := runner.New(t, "go", "run", ".")
	cases := []struct {
		intput, output string
	}{
		{"sample-input-1.txt", "sample-output-1.txt"},
		{"input1.txt", "output1.txt"},
		{"input2.txt", "output2.txt"},
		{"input3.txt", "output3.txt"},
	}
	for _, c := range cases {
		r.Run(c.intput).Want(c.output)
	}
}
