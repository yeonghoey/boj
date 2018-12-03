package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	y, x int
}

type node struct {
	pos
	cost int
}

var offsets = []pos{
	{-2, 1}, {-1, 2}, {1, 2}, {2, 1},
	{2, -1}, {1, -2}, {-1, -2}, {-2, -1},
}

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func nextInt() int {
	scanner.Scan()
	n := 0
	for _, c := range scanner.Bytes() {
		n *= 10
		n += int(c - '0')
	}
	return n
}

func main() {
	scanner.Split(bufio.ScanWords)

	T := nextInt()
	for T > 0 {
		process()
		T--
	}

	_ = writer.Flush()
}

func process() {
	l := nextInt()
	s := pos{nextInt(), nextInt()}
	e := pos{nextInt(), nextInt()}

	v := make([][]bool, l)
	for i := range v {
		v[i] = make([]bool, l)
	}

	v[s.y][s.x] = true
	q := []node{{s, 0}}

	cost := -1
	for len(q) > 0 {
		a := q[0]
		q = q[1:]

		if a.y == e.y && a.x == e.x {
			cost = a.cost
		}

		for _, o := range offsets {
			b := pos{a.y + o.y, a.x + o.x}
			if !(0 <= b.y && b.y < l) || !(0 <= b.x && b.x < l) {
				continue
			}
			if v[b.y][b.x] {
				continue
			}
			q = append(q, node{b, a.cost + 1})
			v[b.y][b.x] = true
		}
	}

	_, _ = fmt.Fprintln(writer, cost)
}
