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
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	N := nextInt()
	G := make([][]int, N)
	for i := range G {
		row := make([]int, 0)
		for j := 0; j < N; j++ {
			e := nextInt()
			if e == 1 {
				row = append(row, j)
			}
		}
		G[i] = row
	}

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	for i := 0; i < N; i++ {
		reachables := bfs(i, N, G)
		line := strings.Trim(fmt.Sprint(reachables), "[]")
		_, _ = fmt.Fprintln(writer, line)
	}
}

func bfs(from, N int, G [][]int) []int {
	reachables := make([]int, N)
	queue := []int{from}
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, y := range G[x] {
			if reachables[y] == 0 {
				reachables[y] = 1
				queue = append(queue, y)
			}
		}
	}
	return reachables
}
