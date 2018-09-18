package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
		{23, 4},
		{32, 1},
		{64, 1},
		{48, 2},
	}

	for _, test := range tests {
		if got := solve(test.x); got != test.want {
			t.Errorf("solve(%d) = %d", test.x, got)
		}
	}
}
