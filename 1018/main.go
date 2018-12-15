package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func nextRow() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func main() {
	N := nextInt()
	M := nextInt()
	board := make([][]byte, N)
	for i := range board {
		board[i] = nextRow()
	}

	best := 8 * 8
	for top := 0; top <= N-8; top++ {
		for left := 0; left <= M-8; left++ {
			x := count(board, top, left)
			best = min(best, x)
		}
	}
	fmt.Println(best)
}

func count(board [][]byte, top, left int) int {
	wb, wbCount := []byte{'W', 'B'}, 0
	for i := top; i < top+8; i++ {
		for j := left; j < left+8; j++ {
			k := (i + j) % 2
			if board[i][j] != wb[k] {
				wbCount++
			}
		}
	}

	bwCount := 8*8 - wbCount
	return min(wbCount, bwCount)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
