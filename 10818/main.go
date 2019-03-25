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
	N := nextInt()
	max := -1000000
	min := +1000000
	for N > 0 {
		x := nextInt()
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
		N--
	}
	fmt.Println(min, max)
}
