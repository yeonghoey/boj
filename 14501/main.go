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

	var f func(int) int
	f = func(d int) int {
		if d > N {
			return 0
		}

		pass := f(d + 1)
		if d+T[d] <= N+1 {
			pick := P[d] + f(d+T[d])
			if pick > pass {
				return pick
			}
		}
		return pass
	}

	answer := f(1)
	fmt.Println(answer)
}
