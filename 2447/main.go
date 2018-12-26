package main

import (
	"bufio"
	"fmt"
	"os"
)

var board [][]byte

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	board = make([][]byte, n)
	for i := range board {
		row := make([]byte, n)
		for j := range row {
			row[j] = ' '
		}
		board[i] = row
	}
	fill(0, 0, n)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	for _, row := range board {
		_, _ = fmt.Fprintln(writer, string(row))
	}
}

func fill(y, x, k int) {
	if k == 1 {
		board[y][x] = '*'
		return
	}

	k3 := k / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			}
			yy := y + i*k3
			xx := x + j*k3
			fill(yy, xx, k3)
		}
	}
}
