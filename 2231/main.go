package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	answer := 0
	for m := 1; m < N; m++ {
		if toN(m) == N {
			answer = m
			break
		}
	}
	fmt.Println(answer)
}

func toN(m int) int {
	x := m
	s := strconv.Itoa(m)
	for _, r := range s {
		x += int(r - '0')
	}
	return x
}
