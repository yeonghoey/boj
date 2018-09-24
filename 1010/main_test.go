package main

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		n, m int
		want int
	}{
		{2, 2, 1},
		{1, 5, 5},
		{13, 29, 67863915},
		{15, 30, 155117520},
		{16, 30, 145422675},
		{17, 30, 119759850},
	}

	for _, c := range cases {
		if got := solve(c.n, c.m); got != c.want {
			t.Errorf("solve(%d, %d) = %d, want %d", c.n, c.m, got, c.want)
		}
	}
}
