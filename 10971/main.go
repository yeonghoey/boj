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
	W := make([][]int, N)
	for i := range W {
		row := make([]int, N)
		for j := range row {
			row[j] = nextInt()
		}
		W[i] = row
	}

	var f func(int, int, int) int
	f = func(k, node, bitmask int) int {
		if k == N-1 {
			return W[node][0]
		}
		best := 0
		for next := 0; next < N; next++ {
			b := 1 << uint(next)
			if bitmask&b > 0 {
				continue
			}
			if W[node][next] == 0 {
				continue
			}
			cost := f(k+1, next, bitmask|b)
			if cost == 0 {
				continue
			}
			cost += W[node][next]
			if best == 0 || cost < best {
				best = cost
			}
		}
		return best
	}
	answer := f(0, 0, 1)
	fmt.Println(answer)
}
