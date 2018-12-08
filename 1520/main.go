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

func main() {
	M := nextInt()
	N := nextInt()
	A := make([][]int, M)
	D := make([][]int, M)
	for i := range A {
		A[i] = make([]int, N)
		D[i] = make([]int, N)
		for j := range A[i] {
			A[i][j] = nextInt()
			D[i][j] = -1
		}
	}

	D[0][0] = 1

	var dp func(y, x int) int

	dp = func(y, x int) int {
		if D[y][x] != -1 {
			return D[y][x]
		}

		moves := []struct{ dy, dx int }{
			{-1, 0}, {0, 1}, {1, 0}, {0, -1},
		}
		total := 0
		for _, m := range moves {
			yy := y + m.dy
			xx := x + m.dx
			if !(0 <= yy && yy < M) || !(0 <= xx && xx < N) {
				continue
			}
			if A[yy][xx] > A[y][x] {
				total += dp(yy, xx)
			}
		}
		D[y][x] = total
		return D[y][x]
	}

	answer := dp(M-1, N-1)
	fmt.Println(answer)
}
