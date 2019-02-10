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

func nextRow() []int {
	scanner.Scan()
	bs := scanner.Bytes()
	ns := make([]int, len(bs))
	for i := range ns {
		ns[i] = int(bs[i] - '0')
	}
	return ns
}

func main() {
	N := nextInt()
	M := nextInt()
	A := make([][]int, N)
	for i := range A {
		A[i] = nextRow()
	}
	B := make([][]int, N)
	for i := range B {
		B[i] = nextRow()
	}

	flip := func(r, c int) bool {
		if r+3 > N || c+3 > M {
			return false
		}
		for i := r; i < r+3; i++ {
			for j := c; j < c+3; j++ {
				A[i][j] = (A[i][j] + 1) % 2
			}
		}
		return true
	}

	try := func() int {
		count := 0
		for i, row := range A {
			for j, value := range row {
				if value != B[i][j] {
					ok := flip(i, j)
					if ok {
						count++
					} else {
						return -1
					}

				}
			}
		}
		return count
	}

	answer := try()
	fmt.Println(answer)
}
