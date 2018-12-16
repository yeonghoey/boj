package main

import (
	"bufio"
	"fmt"
	"os"
)

const empty = 0

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
	M := nextInt()
	A := []int{empty, 1, 0, 0}
	for M > 0 {
		X := nextInt()
		Y := nextInt()
		A[X], A[Y] = A[Y], A[X]
		M--
	}

	for i, a := range A {
		if a == 1 {
			fmt.Println(i)
		}
	}
}
