package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	last := []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	next := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 2; i <= N; i++ {
		next[0] = last[1]
		for d := 1; d <= 8; d++ {
			next[d] = (last[d-1] + last[d+1]) % 1000000000
		}
		next[9] = last[8]

		last, next = next, last
	}

	answer := 0
	for _, x := range last {
		answer += x
		answer %= 1000000000
	}
	fmt.Println(answer)
}
