package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 6, 1},
		{2, 1000, 6},
		{3, 7, 7},
		{6, 2, 6},
		{7, 100, 1},
		{9, 635, 9},
		{10, 71231, 10},
		{12, 6, 4},
		{99, 8, 1},
	}

	for _, test := range tests {
		if got := solve(test.a, test.b); got != test.want {
			t.Errorf("solve(%d, %d) = %d, want: %d",
				test.a, test.b, got, test.want)
		}
	}
}
