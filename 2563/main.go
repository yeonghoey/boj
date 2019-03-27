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
	plane := make([][]bool, 101)
	for i := range plane {
		plane[i] = make([]bool, 101)
	}
	area := 0
	for i := 0; i < N; i++ {
		x := nextInt()
		y := nextInt()
		for xx := x; xx < x+10; xx++ {
			for yy := y; yy < y+10; yy++ {
				if !plane[yy][xx] {
					plane[yy][xx] = true
					area++
				}
			}
		}
	}
	fmt.Println(area)
}
