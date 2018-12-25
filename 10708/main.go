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
	M := nextInt()
	A := make([]int, M)
	for i := range A {
		A[i] = nextInt()
	}

	scores := make([]int, N+1)

	for _, target := range A {
		miss := 0
		for i := 1; i <= N; i++ {
			pick := nextInt()
			if pick == target {
				scores[i]++
			} else {
				miss++
			}
		}
		scores[target] += miss
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= N; i++ {
		_, _ = fmt.Fprintln(writer, scores[i])
	}
	_ = writer.Flush()
}
