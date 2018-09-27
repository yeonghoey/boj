package main

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		as   []int
		want int
	}{
		{[]int{3}, 9},
		{[]int{4, 2}, 8},
		{[]int{5, 7}, 35},
		{[]int{2, 3, 4, 6}, 12},
		{[]int{4, 8, 2, 16, 32}, 64},
		{[]int{4, 8, 2, 16, 32}, 64},
		{[]int{2, 3, 4, 6, 8, 12, 16, 24}, 48},
	}

	for _, c := range cases {
		if got := solve(c.as); got != c.want {
			t.Errorf("solve(%v) = %d, want %d", c.as, got, c.want)
		}
	}
}
