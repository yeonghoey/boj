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
	exists := make([]bool, 42)
	count := 0
	for i := 0; i < 10; i++ {
		x := nextInt()
		m := x % 42
		if !exists[m] {
			exists[m] = true
			count++
		}
	}
	fmt.Println(count)
}
