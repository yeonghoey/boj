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
	seats := make([]bool, 101)
	refused := 0
	for N > 0 {
		x := nextInt()
		if seats[x] {
			refused++
		} else {
			seats[x] = true
		}
		N--
	}
	fmt.Println(refused)
}
