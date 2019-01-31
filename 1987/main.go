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

func nextRow() []int {
	scanner.Scan()
	bs := scanner.Bytes()
	ns := make([]int, len(bs))
	for i := range ns {
		ns[i] = int(bs[i] - 'A')
	}
	return ns
}

func main() {
	R := nextInt()
	C := nextInt()
	board := make([][]int, R)
	for i := range board {
		board[i] = nextRow()
	}

	visited := make([]bool, 26)
	var f func(int, int) int
	f = func(r, c int) int {
		if !(0 <= r && r < R) {
			return 0
		}
		if !(0 <= c && c < C) {
			return 0
		}

		ch := board[r][c]
		if visited[ch] {
			return 0
		}

		visited[ch] = true
		top := f(r-1, c)
		right := f(r, c+1)
		bottom := f(r+1, c)
		left := f(r, c-1)
		visited[ch] = false

		return max(top, right, bottom, left) + 1
	}

	answer := f(0, 0)
	fmt.Println(answer)
}

func max(a0 int, as ...int) int {
	x := a0
	for _, a1 := range as {
		if a1 > x {
			x = a1
		}
	}
	return x
}
