package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	G := make([][]int, N+1)
	for i := 0; i < M; i++ {
		A := nextInt()
		B := nextInt()
		G[B] = append(G[B], A)
	}

	explored := make([]bool, N+1)
	ordered := make([]int, 0, N)

	var dfs func(int)
	dfs = func(x int) {
		explored[x] = true
		for _, y := range G[x] {
			if !explored[y] {
				dfs(y)
			}
		}
		ordered = append(ordered, x)
	}

	for x := 1; x <= N; x++ {
		if !explored[x] {
			dfs(x)
		}
	}

	s := strings.Trim(fmt.Sprint(ordered), "[]")
	fmt.Println(s)
}
