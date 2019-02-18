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
	A := make([]int, N)
	for i := range A {
		A[i] = nextInt()
	}

	result := A[N-1]
	counters := make([]int, 21)
	counters[result] = 1
	for i := N - 2; i >= 1; i-- {
		y := A[i]
		countersNext := make([]int, 21)
		for x, cnt := range counters {
			a := x + y
			if 0 <= a && a <= 20 {
				countersNext[a] += cnt
			}
			b := x - y
			if 0 <= b && b <= 20 {
				countersNext[b] += cnt
			}
		}
		counters = countersNext
	}
	fmt.Println(counters[A[0]])
}
