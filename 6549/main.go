package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	for {
		n := nextInt()
		if n == 0 {
			break
		}
		a := make([]int, n)
		for i := range a {
			a[i] = nextInt()
		}
		answer := solve(a, 0, n)
		_, _ = fmt.Fprintln(writer, answer)
	}
	_ = writer.Flush()
}

func solve(a []int, begin, end int) int {
	if begin >= end {
		return 0
	}
	mid := (begin + end) / 2

	left := solve(a, begin, mid)
	right := solve(a, mid+1, end)
	pivot := a[mid]

	best := max(left, pivot, right)

	i, j := mid, mid
	height := pivot
	for i > begin || j < end-1 {
		ai := -1
		if i-1 >= begin {
			ai = a[i-1]
		}
		aj := -1
		if j+1 < end {
			aj = a[j+1]
		}
		if ai > aj {
			height = min(height, ai)
			i--
		} else {
			height = min(height, aj)
			j++
		}
		best = max(best, height*(j-i+1))
	}
	return best
}

func min(a0 int, as ...int) int {
	x := a0
	for _, a := range as {
		if a < x {
			x = a
		}
	}
	return x
}

func max(a0 int, as ...int) int {
	x := a0
	for _, a := range as {
		if a > x {
			x = a
		}
	}
	return x
}
