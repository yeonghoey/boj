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
	x, r := B/100, B%100
	y, z := r/10, r%10
	fmt.Printf("%d\n%d\n%d\n%d\n", A*z, A*y, A*x, A*B)
}
