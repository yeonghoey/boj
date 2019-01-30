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
		{"editor.in1", "editor.ou1"},
		{"editor.in2", "editor.ou2"},
		{"editor.in3", "editor.ou3"},
		{"editor.in4", "editor.ou4"},
		{"editor.in5", "editor.ou5"},
		{"editor.in6", "editor.ou6"},
		{"editor.in7", "editor.ou7"},
		{"editor.in8", "editor.ou8"},
		{"editor.in9", "editor.ou9"},
		{"editor.ina", "editor.oua"},
	}
	for _, c := range cases {
		r.Run(c.intput).Want(c.output)
	}
}
