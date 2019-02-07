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
	X := nextInt()
	Y := nextInt()
	Z := (100 * Y) / X
	r := (100 * Y) % X
	denom := 99 - Z
	answer := -1
	if denom > 0 {
		numer := X - r
		answer = (X - r) / denom
		if numer%denom > 0 {
			answer++
		}
	}
	fmt.Println(answer)
}
