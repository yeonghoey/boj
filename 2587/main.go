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

func main() {
	A := make([]int, 5)
	S := 0
	for i := range A {
		x := nextInt()
		A[i] = x
		S += x
	}
	sort.Ints(A)
	fmt.Println(S / len(A))
	fmt.Println(A[len(A)/2])
}
