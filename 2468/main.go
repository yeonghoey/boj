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
	sign, x := 1, 0
	for _, b := range scanner.Bytes() {
		if b == '-' {
			sign = -1
			continue
		}
		x *= 10
		x += int(b - '0')
	}
	return sign * x
}

func main() {
	N := nextInt()
	A := make([][]int, N)

	min := 100
	max := 1
	for i := range A {
		row := make([]int, N)
		for j := range row {
			h := nextInt()
			if h < min {
				min = h
			}
			if h > max {
				max = h
			}
			row[j] = h
		}
		A[i] = row
	}

	var mark func(h, r, c int, visited [][]bool)
	mark = func(h, r, c int, visited [][]bool) {
		if !(0 <= r && r < N) || !(0 <= c && c < N) {
			return
		}
		if !(A[r][c] > h) {
			return
		}
		if visited[r][c] {
			return
		}
		visited[r][c] = true
		mark(h, r-1, c, visited)
		mark(h, r+1, c, visited)
		mark(h, r, c-1, visited)
		mark(h, r, c+1, visited)
	}

	try := func(h int) int {
		count := 0

		visited := make([][]bool, N)
		for i := range visited {
			visited[i] = make([]bool, N)
		}

		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if A[i][j] > h && !visited[i][j] {
					mark(h, i, j, visited)
					count++
				}
			}
		}
		return count
	}

	// At least 1, if it doesn't rain
	best := 1
	for h := min; h < max; h++ {
		count := try(h)
		if count > best {
			best = count
		}
	}
	fmt.Println(best)
}
