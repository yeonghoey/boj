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
	T := nextInt()
	for T > 0 {
		main2()
		T--
	}
}

func main2() {
	N := nextInt()
	G := make([][]int, N+1)
	for tail := 1; tail <= N; tail++ {
		head := nextInt()
		G[tail] = append(G[tail], head)
	}

	explored := make([]bool, N+1)

	var detectCycle func(start, this int) bool
	detectCycle = func(start, this int) bool {
		explored[this] = true
		for _, next := range G[this] {
			if next == start {
				return true
			}

			if !explored[next] {
				if detectCycle(start, next) {
					return true
				}
			}
		}
		return false
	}

	numCycles := 0
	for start := 1; start <= N; start++ {
		if !explored[start] {
			if detectCycle(start, start) {
				numCycles++
			}
		}
	}

	fmt.Println(numCycles)
}
