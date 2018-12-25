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

type pos struct {
	y, x int
}

func main() {
	board := make([][]int, 9)
	blanks := make([]pos, 0)
	for i := range board {
		board[i] = make([]int, 9)
		for j := range board[i] {
			x := nextInt()
			board[i][j] = x
			if x == 0 {
				blanks = append(blanks, pos{i, j})
			}
		}
	}

	pickCandidates := func(p pos) []int {
		exists := make([]bool, 9+1)
		for i := 0; i < 9; i++ {
			exists[board[p.y][i]] = true
			exists[board[i][p.x]] = true
		}

		lty := (p.y / 3) * 3
		ltx := (p.x / 3) * 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				exists[board[lty+i][ltx+j]] = true
			}
		}

		candidates := make([]int, 0)
		for i := 1; i <= 9; i++ {
			if !exists[i] {
				candidates = append(candidates, i)
			}
		}

		return candidates
	}

	var walk func(int) bool
	walk = func(bi int) bool {
		if bi == len(blanks) {
			return true
		}
		p := blanks[bi]
		for _, cand := range pickCandidates(p) {
			board[p.y][p.x] = cand
			if walk(bi + 1) {
				return true
			}
			board[p.y][p.x] = 0
		}
		// Unexpected
		return false
	}
	walk(0)
	writer := bufio.NewWriter(os.Stdout)
	for _, row := range board {
		_, _ = fmt.Fprint(writer, row[0])
		for i := 1; i < 9; i++ {
			_, _ = fmt.Fprintf(writer, " %d", row[i])
		}
		_, _ = fmt.Fprintln(writer)
	}
	_ = writer.Flush()
}
