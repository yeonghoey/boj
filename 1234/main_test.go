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
		{"input1.txt", "output1.txt"},
	}
	for _, s := range samples {
		r.Run(s.intput).Want(s.output)
	}
}
