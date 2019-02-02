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
	T := make([][]int, N)
	S := []int{}
	for i := 0; i < N; i++ {
		p := nextInt()
		if p == -1 {
			S = append(S, i)
			continue
		}
		T[p] = append(T[p], i)
	}

	X := nextInt()

	var f func(int) int
	f = func(x int) int {
		if x == X {
			return 0
		}

		count := 0
		for _, c := range T[x] {
			count += f(c)
		}

		return max(count, 1)
	}

	answer := 0
	for _, s := range S {
		answer += f(s)
	}
	fmt.Println(answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
