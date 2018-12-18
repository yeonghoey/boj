package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		line(n, i)
	}
	for i := n - 1; i >= 1; i-- {
		line(n, i)
	}
}

func line(n, i int) {
	put(" ", n-i)
	put("*", i*2-1)
	put("\n", 1)
}

func put(s string, count int) {
	ss := strings.Repeat(s, count)
	fmt.Print(ss)
}
