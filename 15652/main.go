package main

import (
	"bufio"
	"bytes"
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
	N := nextInt()
	M := nextInt()

	A := make([]int, M+1)

	var buf bytes.Buffer

	var f func(int, int)
	f = func(n, m int) {
		if m > M {
			for i := 1; i <= M; i++ {
				buf.WriteByte(byte(A[i] + '0'))
				if i < M {
					buf.WriteByte(' ')
				}
			}
			buf.WriteByte('\n')
			return
		}
		for x := n; x <= N; x++ {
			A[m] = x
			f(x, m+1)
		}
	}
	f(1, 1)
	fmt.Println(buf.String())
}
