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
	M := nextInt()
	m := make([]int, N)
	for i := range m {
		m[i] = nextInt()
	}
	c := make([]int, N)
	totalCost := 0
	for i := range c {
		c[i] = nextInt()
		totalCost += c[i]
	}

	prev := make([]int, totalCost+1)
	this := make([]int, totalCost+1)
	for i := 0; i < N; i++ {
		for cost := range prev {
			this[cost] = prev[cost]
			cp := cost - c[i]
			if cp < 0 {
				continue
			}
			x := prev[cp] + m[i]
			if x > this[cost] {
				this[cost] = x
			}
		}
		prev, this = this, prev
	}
	var answer int
	for cost, mx := range prev {
		if mx >= M {
			answer = cost
			break
		}
	}
	fmt.Println(answer)
}
