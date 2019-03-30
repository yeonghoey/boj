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
	A := nextInt()
	B := nextInt()
	C := nextInt()

	if A > B {
		A, B = B, A
	}
	if A > C {
		A, C = C, A
	}
	if B > C {
		B, C = C, B
	}

	var answer int
	x := B - A - 1
	y := C - B - 1
	if x > y {
		answer = x
	} else {
		answer = y
	}
	fmt.Println(answer)
}
