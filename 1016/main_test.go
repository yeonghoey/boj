package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var min, max int64
	for min = 1; min < 100; min++ {
		for max = min + 1; max <= 100; max++ {
			want := solveBrute(min, max)
			got := solve(min, max)
			if got != want {
				t.Errorf("solve(%d, %d) = %d, want %d", min, max, got, want)
			}
		}
	}
}

func solveBrute(min, max int64) int {
	count := 0
	for i := min; i <= max; i++ {
		found := false
		var j int64
		for j = 2; j <= i; j++ {
			if i%(j*j) == 0 {
				found = true
				break
			}
		}

		if !found {
			count++
		}
	}
	return count
}
