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
	var mx, mi int
	for i := 1; i <= 9; i++ {
		x := nextInt()
		if x > mx {
			mx = x
			mi = i
		}
	}
	fmt.Printf("%d\n%d\n", mx, mi)
}
