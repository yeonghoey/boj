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
	fmt.Println(kth())
}

func kth() int {
	N := nextInt()
	K := nextInt()
	deleted := make([]bool, N+1)
	k := 0
	for x := 2; x <= N; x++ {
		if !deleted[x] {
			for y := x; y <= N; y += x {
				if !deleted[y] {
					deleted[y] = true
					if k++; k == K {
						return y
					}
				}
			}
		}
	}
	return -1
}
