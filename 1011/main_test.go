package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{0, 1, 1},
		{0, 2, 2},
		{0, 3, 3},
		{0, 4, 3},
		{0, 5, 4},
		{0, 6, 4},
		{0, 7, 5},
		{0, 8, 5},
		{0, 9, 5},
		{0, 10, 6},
	}

	for _, test := range tests {
		if got := solve(test.x, test.y); got != test.want {
			t.Errorf("solve(%d, %d) = %d, want %d", test.x, test.y, got, test.want)
		}
	}
}
