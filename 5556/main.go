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
	writer := bufio.NewWriter(os.Stdout)
	N := nextInt()
	K := nextInt()
	for K > 0 {
		a := nextInt()
		b := nextInt()
		x := min(a, b, N-a+1, N-b+1)
		color := (x-1)%3 + 1
		_, _ = fmt.Fprintln(writer, color)
		K--
	}
	_ = writer.Flush()
}

func min(a0 int, as ...int) int {
	x := a0
	for _, a1 := range as {
		if a1 < x {
			x = a1
		}
	}
	return x
}
