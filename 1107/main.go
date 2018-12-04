package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	x := 0
	for _, c := range scanner.Bytes() {
		x *= 10
		x += int(c - '0')
	}
	return x
}

func init() {
	scanner.Split(bufio.ScanWords)
}

func main() {
	N := nextInt()
	M := nextInt()
	malfuncs := make([]bool, 10)
	for i := 0; i < M; i++ {
		b := nextInt()
		malfuncs[b] = true
	}

	answer := abs(N - 100)
	if M < 10 {
		x := closest(N, M, malfuncs)
		answer = min(answer, len(strconv.Itoa(x))+abs(N-x))
	}
	fmt.Println(answer)
}

func closest(N, M int, malfuncs []bool) int {
	down, up := N, N
	for {
		// Test 'down' first!
		if down >= 0 && check(down, malfuncs) {
			return down
		}
		if check(up, malfuncs) {
			return up
		}
		down--
		up++
	}
}

func check(x int, malfuncs []bool) bool {
	s := strconv.Itoa(x)
	for _, d := range s {
		b := int(d - '0')
		if malfuncs[b] {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
