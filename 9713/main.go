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
	w := bufio.NewWriter(os.Stdout)
	T := nextInt()
	for T > 0 {
		N := nextInt()
		sum := 0
		for i := 1; i <= N; i += 2 {
			sum += i
		}
		_, _ = fmt.Fprintln(w, sum)
		T--
	}
	_ = w.Flush()
}
