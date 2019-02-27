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
	M := nextInt()
	A := make([]int, M+1)
	for i := 0; i < N; i++ {
		for j := 1; j <= M; j++ {
			if A[j-1] > A[j] {
				A[j] = A[j-1]
			}
			A[j] += nextInt()
		}
	}
	fmt.Println(A[M])
}
