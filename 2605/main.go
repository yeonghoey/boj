package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	N := nextInt()
	A := make([]int, N)
	for i := range A {
		A[i] = i + 1
		x := nextInt()
		for j := i; j > i-x; j-- {
			A[j], A[j-1] = A[j-1], A[j]
		}
	}
	s := fmt.Sprint(A)
	s = strings.Trim(s, "[]")
	fmt.Println(s)
}
