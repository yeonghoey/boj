package main

import (
	"testing"

	"github.com/yeonghoey/boj/runner"
)

func TestSamples(t *testing.T) {
	r := runner.New(t, "go", "run", ".")
	samples := []struct {
		intput, output string
	}{
		{"sample-input-1.txt", "sample-output-1.txt"},
		{"sample-input-2.txt", "sample-output-2.txt"},
		{"sample-input-3.txt", "sample-output-3.txt"},
	}
	for _, s := range samples {
		r.Run(s.intput).Want(s.output)
	}
}