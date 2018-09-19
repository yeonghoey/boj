package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		smap []string
		want int
	}{
		{
			[]string{
				"1100000000",
				"0100000000",
				"0000100000",
				"0000100000",
				"0011000111",
				"0000100111",
				"0000000111",
				"0000000000",
			},
			5,
		},
		{
			[]string{
				"00",
				"00",
			},
			0,
		},
		{
			[]string{
				"10",
				"00",
			},
			1,
		},
		{
			[]string{
				"10",
				"01",
			},
			2,
		},
		{
			[]string{
				"11",
				"01",
			},
			1,
		},
		{
			[]string{
				"01",
				"11",
			},
			1,
		},
		{
			[]string{
				"011",
				"101",
				"111",
			},
			1,
		},
	}

	for _, test := range tests {
		m, n, _, poss := parse(test.smap)
		got := solve(m, n, poss)
		if got != test.want {
			t.Errorf("solve(%v) = %d", test.smap, got)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		smap    []string
		m, n, k int
		poss    map[pos]bool
	}{
		{
			smap: []string{
				"1100000000",
				"0100000000",
				"0000100000",
				"0000100000",
				"0011000111",
				"0000100111",
				"0000000111",
				"0000000000",
			},
			m: 10,
			n: 8,
			k: 17,
			poss: map[pos]bool{
				pos{0, 0}: true, pos{1, 0}: true, pos{1, 1}: true, pos{4, 2}: true, pos{4, 3}: true,
				pos{4, 5}: true, pos{2, 4}: true, pos{3, 4}: true, pos{7, 4}: true, pos{8, 4}: true,
				pos{9, 4}: true, pos{7, 5}: true, pos{8, 5}: true, pos{9, 5}: true, pos{7, 6}: true,
				pos{8, 6}: true, pos{9, 6}: true,
			},
		},
	}

	for _, test := range tests {
		m, n, k, poss := parse(test.smap)
		if !(m == test.m && n == test.n && k == test.k) {
			t.Errorf("parse(%v) = (m=%d, n=%d, k=%d)", test.smap, m, n, k)
		}

		for p := range test.poss {
			if !poss[p] {
				t.Errorf("parse(%v) = %v", test.smap, poss)
			}
		}
	}
}

func parse(smap []string) (m, n, k int, poss map[pos]bool) {
	m = len(smap[0])
	n = len(smap)
	k = 0
	poss = make(map[pos]bool)
	for i, row := range smap {
		for j, cell := range row {
			if cell == '1' {
				k++
				poss[pos{j, i}] = true
			}
		}
	}
	return
}
