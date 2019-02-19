package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	heap := make([]int, n+1)
	heap[1] = 1
	l := 1
	for i := 2; i <= n; i++ {
		// Reverse heapify
		k := l
		for k > 1 {
			heap[k] = heap[k/2]
			k /= 2
		}
		l++
		heap[1] = i
		heap[l] = 1
	}

	s := fmt.Sprint(heap[1:])
	s = strings.Trim(s, "[]")
	fmt.Println(s)
}
