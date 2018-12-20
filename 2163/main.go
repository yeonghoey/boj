package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Scanner

func init() {
	in = bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
}

func nextInt() int {
	in.Scan()
	x := 0
	for _, b := range in.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	N := nextInt()
	M := nextInt()

	D := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		D[i] = make([]int, M+1)
		for j := 1; j <= M; j++ {
			D[i][j] = -1
		}
	}
	D[1][1] = 0

	var cut func(int, int) int
	cut = func(n, m int) int {
		if D[n][m] != -1 {
			return D[n][m]
		}

		x := 1
		if n > m {
			x += cut(n/2, m)
			x += cut(n-n/2, m)
		} else {
			x += cut(n, m/2)
			x += cut(n, m-m/2)
		}
		D[n][m] = x
		return x

	}
	answer := cut(N, M)
	fmt.Println(answer)
}
