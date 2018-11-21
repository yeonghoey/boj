package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func nextInt() int {
	var n int
	_, _ = fmt.Fscan(reader, &n)
	return n
}

func main() {
	N := nextInt()
	M := nextInt()

	G := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		G[i] = []int{}
	}

	for i := 0; i < M; i++ {
		a := nextInt()
		b := nextInt()
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	explored := map[int]bool{1: true}
	queue := []int{1}
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, y := range G[x] {
			if !explored[y] {
				explored[y] = true
				queue = append(queue, y)
			}
		}
	}

	// Exclude `1`
	answer := len(explored) - 1
	fmt.Println(answer)
}
