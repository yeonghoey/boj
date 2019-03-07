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
	n := nextInt()
	a := make([]int, n+1)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		for j := 0; j < i; j++ {
			if a[i] > a[j] && d[i] < d[j] {
				d[i] = d[j]
			}
		}
		d[i]++
	}
	answer := 0
	for _, x := range d {
		if x > answer {
			answer = x
		}
	}
	fmt.Println(answer)
}
