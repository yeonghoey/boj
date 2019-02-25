package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	G := make([][]int, N)
	for i := 0; i < M; i++ {
		A := nextInt() - 1
		B := nextInt() - 1
		G[B] = append(G[B], A)
	}

	bfs := func(start int) int {
		count := 1
		explored := make([]bool, N)
		explored[start] = true
		queue := []int{start}
		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]
			for _, y := range G[x] {
				if !explored[y] {
					count++
					explored[y] = true
					queue = append(queue, y)
				}
			}
		}
		return count
	}

	best := 0
	answers := []int{}
	for i := 0; i < N; i++ {
		x := bfs(i)
		if x > best {
			best = x
			answers = []int{i + 1}
		} else if x == best {
			answers = append(answers, i+1)
		}
	}
	sort.Ints(answers)
	line := fmt.Sprint(answers)
	line = strings.Trim(line, "[]")
	fmt.Println(line)
}
