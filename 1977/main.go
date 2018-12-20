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
	M := nextInt()
	N := nextInt()
	sum := 0
	min := -1
	for x := 1; x <= 100; x++ {
		xx := x * x
		if xx >= M && xx <= N {
			sum += xx
			if min == -1 {
				min = xx
			}
		}
	}
	if min != -1 {
		fmt.Println(sum)
	}
	fmt.Println(min)
}
