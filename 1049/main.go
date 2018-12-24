package main

import (
	"bufio"
	"fmt"
	"os"
)

const undefined = -1

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
	M := nextInt()
	D := make([]int, N+1)
	for i := range D {
		D[i] = undefined
	}
	D[0] = 0
	for i := 0; i < M; i++ {
		pack6 := nextInt()
		pack1 := nextInt()
		for k := 1; k <= N; k++ {
			prev6 := max(0, k-6)
			prev1 := max(0, k-1)
			best := min(D[prev6]+pack6, D[prev1]+pack1)

			if D[k] != undefined {
				best = min(best, D[k])
			}
			D[k] = best
		}
	}

	fmt.Println(D[N])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
