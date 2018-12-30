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
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	N := nextInt()
	A := make([][]int, N+1)
	for i := 0; i < N; i++ {
		a := nextInt()
		b := nextInt()
		A[a] = append(A[a], b)
		A[b] = append(A[b], a)
	}

	parents := make([]int, N+1)
	parents[1] = -1
	queue := []int{1}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, other := range A[node] {
			if parents[other] != 0 {
				continue
			}
			parents[other] = node
			queue = append(queue, other)
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	for i := 2; i <= N; i++ {
		_, _ = fmt.Fprintln(writer, parents[i])
	}
}
