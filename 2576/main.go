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
	sum := 0
	min := 100
	for i := 0; i < 7; i++ {
		x := nextInt()
		if x%2 == 1 {
			sum += x
			if x < min {
				min = x
			}
		}
	}
	if sum == 0 && min == 100 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%d\n%d\n", sum, min)
	}
}
