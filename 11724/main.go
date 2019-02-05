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
	G := make([][]int, N+1)
	for i := 0; i < M; i++ {
		a := nextInt()
		b := nextInt()
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	explored := make([]bool, N+1)
	bfs := func(start int) {
		q := []int{start}
		explored[start] = true
		for len(q) > 0 {
			x := q[0]
			q = q[1:]
			for _, y := range G[x] {
				if !explored[y] {
					explored[y] = true
					q = append(q, y)
				}
			}
		}
	}

	count := 0
	for v := 1; v <= N; v++ {
		if !explored[v] {
			bfs(v)
			count++
		}
	}

	fmt.Println(count)
}
