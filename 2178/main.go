package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func nextInt() int {
	var n int
	_, _ = fmt.Fscan(reader, &n)
	return n
}

func nextString() string {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	return s
}

type pos struct {
	y, x int
}

type entry struct {
	pos
	d int
}

var offsets = [...]pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	N := nextInt()
	M := nextInt()
	maze := make([][]int, N)
	for i := range maze {
		maze[i] = make([]int, M)
		row := nextString()
		for j, ch := range row {
			maze[i][j] = int(ch - '0')
		}
	}

	s := pos{0, 0}
	q := make([]entry, 0, N*M)
	q = append(q, entry{s, 1})
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p.x == M-1 && p.y == N-1 {
			fmt.Println(p.d)
			break
		}

		for _, o := range offsets {
			y := p.y + o.y
			x := p.x + o.x

			if !(0 <= x && x < M && 0 <= y && y < N) {
				continue
			}

			if maze[y][x] == 1 {
				maze[y][x] = -1
				q = append(q, entry{pos{y, x}, p.d + 1})
			}
		}
	}
}
