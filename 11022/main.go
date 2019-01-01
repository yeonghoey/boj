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
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	T := nextInt()
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	for i := 1; i <= T; i++ {
		A := nextInt()
		B := nextInt()
		_, _ = fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i, A, B, A+B)
	}
}
