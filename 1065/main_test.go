package main

import "testing"

func TestSolve(t *testing.T) {
	for i := 1; i <= 1000; i++ {
		want := solveBrute(i)
		got := solve(i)
		if got != want {
			t.Errorf("solve(%d) = %d, want: %d", i, got, want)
		}
	}
}

func solveBrute(n int) int {
	count := 0
	for x := 1; x <= n; x++ {
		if x < 100 {
			count++
			continue
		}
		if x == 1000 {
			continue
		}

		a := x / 100
		r := x % 100
		b := r / 10
		c := r % 10

		if (a - b) == (b - c) {
			count++
		}
	}
	return count
}
