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
	S := nextInt()
	A := make([]int, N)
	for i := range A {
		A[i] = nextInt()
	}

	count := 0
	var f func(int, int, int)
	f = func(k, s, picked int) {
		if k == N {
			if picked > 0 && s == S {
				count++
			}
			return
		}
		f(k+1, s, picked)
		f(k+1, s+A[k], picked+1)
	}
	f(0, 0, 0)
	fmt.Println(count)
}
