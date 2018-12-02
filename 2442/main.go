package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		put(" ", n-i-1)
		put("*", i*2+1)
		put("\n", 1)
	}
}

func put(s string, k int) {
	ss := strings.Repeat(s, k)
	fmt.Print(ss)
}
