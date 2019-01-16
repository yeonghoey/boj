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
	M := nextInt()

	A := make([]int, M+1)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var f func(int)
	f = func(m int) {
		if m > M {
			s := strings.Trim(fmt.Sprint(A[1:]), "[]")
			fmt.Fprintln(writer, s)
			return
		}
		for x := 1; x <= N; x++ {
			A[m] = x
			f(m + 1)
		}
	}
	f(1)
}
