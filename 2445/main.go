package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		line(i, n)
	}
	for i := n - 1; i >= 1; i-- {
		line(i, n)
	}
}

func line(x, n int) {
	s := strings.Repeat("*", x)
	b := strings.Repeat(" ", n-x)
	fmt.Printf("%s%s%s%s\n", s, b, b, s)
}
