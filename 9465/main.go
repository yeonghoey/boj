package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func main() {
	in.Split(bufio.ScanWords)
	T := nextInt()
	for T > 0 {
		process()
		T--
	}
}

func process() {
	n := nextInt()
	a := make([][]int, 2)
	a[0] = make([]int, n)
	for i := range a[0] {
		a[0][i] = nextInt()
	}
	a[1] = make([]int, n)
	for i := range a[1] {
		a[1][i] = nextInt()
	}

	d := make([][]int, 2)
	d[0] = make([]int, 3)
	d[1] = make([]int, 3)
	d[0][2] = a[0][0]
	d[1][2] = a[1][0]
	for i := 1; i < n; i++ {
		copy(d[0][:2], d[0][1:])
		copy(d[1][:2], d[1][1:])
		prev2 := max(d[0][0], d[1][0])
		d[0][2] = a[0][i] + max(prev2, d[1][1])
		d[1][2] = a[1][i] + max(prev2, d[0][1])
	}
	fmt.Println(max(d[0][2], d[1][2]))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
