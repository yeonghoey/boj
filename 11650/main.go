package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type coord struct {
	x, y int
}

func main() {
	N := nextInt()
	A := make([]coord, N)
	for i := range A {
		A[i].x = nextInt()
		A[i].y = nextInt()
	}
	sort.Slice(A, func(i, j int) bool {
		if A[i].x == A[j].x {
			return A[i].y < A[j].y
		}
		return A[i].x < A[j].x
	})
	writer := bufio.NewWriter(os.Stdout)
	for _, c := range A {
		_, _ = fmt.Fprintf(writer, "%d %d\n", c.x, c.y)
	}
	_ = writer.Flush()
}
