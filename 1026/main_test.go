package main

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		a, b []int
		want int
	}{
		{
			[]int{1},
			[]int{9},
			9,
		},
		{
			[]int{1, 2},
			[]int{0, 1},
			1,
		},
		{
			[]int{1, 2, 3},
			[]int{2, 1, 5},
			12,
		},
		{
			[]int{1, 1, 1, 6, 0},
			[]int{2, 7, 8, 3, 1},
			18,
		},
	}

	for _, c := range cases {
		if got := solve(c.a, c.b); got != c.want {
			t.Errorf("solve(%v, %v) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
