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
	R := nextInt()
	G := nextInt()
	B := nextInt()

	F := make([]int, 11)
	F[1] = 1
	for i := 2; i <= 10; i++ {
		F[i] = F[i-1] * i
	}

	var f func(int, int, int, int) int
	f = func(lv, r, g, b int) int {
		if r < 0 || g < 0 || b < 0 {
			return 0
		}
		if lv < 1 {
			return 1
		}
		count := 0 +
			f(lv-1, r-lv, g, b) +
			f(lv-1, r, g-lv, b) +
			f(lv-1, r, g, b-lv)
		if lv%2 == 0 {
			h := lv / 2
			x := 0 +
				f(lv-1, r-h, g-h, b) +
				f(lv-1, r, g-h, b-h) +
				f(lv-1, r-h, g, b-h)
			count += x * (F[lv] / F[h] / F[h])
		}
		if lv%3 == 0 {
			t := lv / 3
			x := f(lv-1, r-t, g-t, b-t)
			count += x * (F[lv] / F[t] / F[t] / F[t])
		}
		return count
	}

	answer := f(N, R, G, B)
	fmt.Println(answer)
}
