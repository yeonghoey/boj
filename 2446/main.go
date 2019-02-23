package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	put := func(s string, count int) {
		ss := strings.Repeat(s, count)
		_, _ = fmt.Fprint(w, ss)
	}

	for i := N - 1; i > 0; i-- {
		put(" ", N-i-1)
		put("*", i*2+1)
		put("\n", 1)
	}
	for i := 0; i < N; i++ {
		put(" ", N-i-1)
		put("*", i*2+1)
		put("\n", 1)
	}
}
