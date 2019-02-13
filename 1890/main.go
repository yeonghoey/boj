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
	for i := range A {
		A[i] = make([]int, N)
	}
	A[0][0] = 1
	for i := range A {
		for j := range A[i] {
			x := nextInt()
			if i == N-1 && j == N-1 {
				continue
			}
			if i+x < N {
				A[i+x][j] += A[i][j]
			}
			if j+x < N {
				A[i][j+x] += A[i][j]
			}
		}
	}
	fmt.Println(A[N-1][N-1])
}
