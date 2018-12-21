package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	for i := n; i >= 1; i-- {
		put(" ", n-i)
		put("*", i*2-1)
		put("\n", 1)
	}
}

func put(s string, n int) {
	fmt.Print(strings.Repeat(s, n))
}
