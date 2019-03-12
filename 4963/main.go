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
	w := bufio.NewWriter(os.Stdout)
	for {
		x := sub()
		if x < 0 {
			break
		}
		_, _ = fmt.Fprintln(w, x)
	}
	_ = w.Flush()
}

func sub() int {
	w := nextInt()
	h := nextInt()
	if w == 0 && h == 0 {
		return -1
	}
	M := make([][]int, h)
	for i := range M {
		row := make([]int, w)
		for j := range row {
			row[j] = nextInt()
		}
		M[i] = row
	}
	var f func(r, c int)
	f = func(r, c int) {
		for _, offset := range []struct{ r, c int }{
			{-1, 0}, {-1, 1}, {0, 1}, {1, 1},
			{1, 0}, {1, -1}, {0, -1}, {-1, -1},
		} {
			rr := r + offset.r
			cc := c + offset.c
			if !(0 <= rr && rr < h) || !(0 <= cc && cc < w) {
				continue
			}
			if M[rr][cc] != 1 {
				continue
			}
			M[rr][cc] = -1
			f(rr, cc)
		}
	}
	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if M[i][j] == 1 {
				f(i, j)
				count++
			}
		}
	}
	return count
}
