package main

import (
	"testing"

	"github.com/yeonghoey/boj/runner"
)

func TestSamples(t *testing.T) {
	r := runner.New(t, "go", "run", ".")
	r.Run("input1.txt").Want("output1.txt")
	r.Run("input2.txt").Want("output2.txt")
	r.Run("input3.txt").Want("output3.txt")
}
