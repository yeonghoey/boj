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
		{"sample-input-2.txt", "sample-output-2.txt"},
		{"sample-input-3.txt", "sample-output-3.txt"},
		{"custom-input-1.txt", "custom-output-1.txt"},
		{"custom-input-2.txt", "custom-output-2.txt"},
		{"custom-input-3.txt", "custom-output-3.txt"},
	}
	for _, c := range cases {
		r.Run(c.intput).Want(c.output)
	}
}
