package main

import (
	"bufio"
	"fmt"
	"os"
)

type revision struct {
	from int
	diff int
}

var scanner *bufio.Scanner
var writer *bufio.Writer

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
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
	K := nextInt()

	A := make([]int, N+1)
	S := make([]int, N+1)
	sum := 0
	for i := 1; i <= N; i++ {
		a := nextInt()
		A[i] = a
		sum += a
		S[i] = sum
	}

	revs := make([]revision, 0, M)
	for i := 0; i < M+K; i++ {
		a := nextInt()
		b := nextInt()
		c := nextInt()

		// Revise
		if a == 1 {
			revs = append(revs, revision{b, c - A[b]})
			A[b] = c
		}

		// Output
		if a == 2 {
			lo := S[b-1]
			hi := S[c]
			for _, r := range revs {
				if b-1 >= r.from {
					lo += r.diff
				}
				if c >= r.from {
					hi += r.diff
				}
			}

			_, _ = fmt.Fprintln(writer, hi-lo)
		}
	}

	_ = writer.Flush()
}
