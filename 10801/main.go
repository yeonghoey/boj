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

func main() {
	A := make([]int, 10)
	B := make([]int, 10)
	for i := range A {
		A[i] = nextInt()
	}
	for i := range B {
		B[i] = nextInt()
	}

	winA := 0
	winB := 0
	for i := 0; i < 10; i++ {
		a := A[i]
		b := B[i]
		if a > b {
			winA++
		}
		if a < b {
			winB++
		}
	}
	winner := 'D'
	if winA > winB {
		winner = 'A'
	}
	if winA < winB {
		winner = 'B'
	}
	fmt.Printf("%c\n", winner)
}
