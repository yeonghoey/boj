package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pos struct{ r, c int }

func main() {
	grid := input()
	subdivs := make([]int, 0)
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				subdivs = append(subdivs, count(grid, i, j))
			}
		}
	}

	output(subdivs)
}

func input() [][]byte {
	var N int
	_, _ = fmt.Fscanf(reader, "%d\n", &N)
	grid := make([][]byte, N)
	for i := range grid {
		grid[i] = make([]byte, N)
		for j := range grid[i] {
			grid[i][j], _ = reader.ReadByte()
			grid[i][j] -= '0'
		}
		_, _ = reader.ReadByte()
	}
	return grid
}

func count(grid [][]byte, row, col int) int {
	n := len(grid)
	k := 0
	q := []pos{{row, col}}
	for len(q) > 0 {
		r, c := q[0].r, q[0].c
		q = q[1:]

		if !((0 <= r && r < n) && (0 <= c && c < n)) {
			continue
		}

		if grid[r][c] != 1 {
			continue
		}
		grid[r][c] = 0
		k++

		q = append(q,
			pos{r - 1, c},
			pos{r, c + 1},
			pos{r + 1, c},
			pos{r, c - 1},
		)
	}
	return k
}

func output(subdivs []int) {
	sort.Ints(subdivs)
	fmt.Fprintln(writer, len(subdivs))
	for _, k := range subdivs {
		fmt.Fprintln(writer, k)
	}
	_ = writer.Flush()
}
