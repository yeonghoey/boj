package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	maxNumber = 10000
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
	N := nextInt()
	a := make([]int, maxNumber+1)
	for i := 0; i < N; i++ {
		a[nextInt()]++
	}

	for i := 1; i <= maxNumber; i++ {
		for j := 0; j < a[i]; j++ {
			_, _ = fmt.Fprintln(writer, i)
		}
	}
	_ = writer.Flush()
}
