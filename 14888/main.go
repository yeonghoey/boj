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
	A := make([]int, N)
	for i := range A {
		A[i] = nextInt()
	}
	add := nextInt()
	sub := nextInt()
	mul := nextInt()
	div := nextInt()

	min := +1000000000
	max := -1000000000
	var f func(x, k int)
	f = func(x, k int) {
		if k == N {
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
			return
		}
		y := A[k]
		if add > 0 {
			add--
			f(x+y, k+1)
			add++
		}
		if sub > 0 {
			sub--
			f(x-y, k+1)
			sub++
		}
		if mul > 0 {
			mul--
			f(x*y, k+1)
			mul++
		}
		if div > 0 {
			div--
			f(x/y, k+1)
			div++
		}
	}
	f(A[0], 1)
	fmt.Println(max)
	fmt.Println(min)
}
