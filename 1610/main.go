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
	C := nextInt()
	T := nextInt()
	cost := make([]int, N)
	time := make([]int, N)
	for i := 1; i < N; i++ {
		cost[i] = nextInt()
		time[i] = nextInt()
	}

	cSum := 0
	cumCost := make([]int, N)
	tSum := 0
	cumTime := make([]int, N)
	for i := N - 1; i >= 1; i-- {
		cSum += cost[i]
		cumCost[i] = cSum
		tSum += time[i]
		cumTime[i] = tSum
	}

	remain := C
	var f func(u, t int) int
	f = func(u, t int) int {
		if u == N-1 {
			maxByCost := remain / cost[u]
			maxByTime := (T - t) / time[u]
			count := max(0, min(maxByCost, maxByTime))
			remain -= count * cost[u]
			return count
		}
		now := t
		count := 0
		for {
			if now+cumTime[u] > T {
				break
			}
			if remain < cumCost[u] {
				break
			}
			now += time[u]
			remain -= cost[u]
			count += f(u+1, now)
		}
		return count
	}
	answer := 1
	if N > 1 {
		answer = f(1, 0)
	}
	fmt.Println(answer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
