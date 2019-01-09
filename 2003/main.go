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
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = A[i-1] + nextInt()
	}

	count := 0
	lo, hi := 0, 0
	for lo <= N && hi <= N {
		sum := A[hi] - A[lo]

		if sum == M {
			count++
		}

		if sum <= M {
			hi++
		} else {
			lo++
		}
	}

	fmt.Println(count)
}
