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
		{"i1", "o1"},
		{"i2", "o2"},
	}
	for _, c := range cases {
		r.Run(c.intput).Want(c.output)
	}
}
