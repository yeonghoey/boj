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
	T := make([]int, N+1)
	P := make([]int, N+1)
	for d := 1; d <= N; d++ {
		T[d] = nextInt()
		P[d] = nextInt()
	}

	D := make([]int, N+2)
	for d := N; d >= 1; d-- {
		D[d] = D[d+1]
		next := d + T[d]
		if next <= N+1 {
			pick := P[d] + D[next]
			if pick > D[d] {
				D[d] = pick
			}
		}
	}

	answer := D[1]
	fmt.Println(answer)
}
