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
	N := nextInt()
	A := make([]int, N+1)
	D := make([][]int, N+1)

	for i := 0; i <= N; i++ {
		if i > 0 {
			A[i] = nextInt()
		}
		D[i] = make([]int, N+1)
	}

	for ai := 1; ai <= N; ai++ {
		for ni := 1; ni < ai; ni++ {
			D[ai][ni] = D[ai-1][ni]
		}
		for ni := ai; ni <= N; ni++ {
			D[ai][ni] = max(D[ai-1][ni], D[ai][ni-ai]+A[ai])
		}
	}

	fmt.Println(D[N][N])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
