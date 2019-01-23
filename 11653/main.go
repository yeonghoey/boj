package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int
	_, _ = fmt.Scan(&x)

	factors := make([]int, 0)
	f := func(p int) {
		for x%p == 0 {
			factors = append(factors, p)
			x /= p
		}
	}

	f(2)
	for p := 3; p*p <= x; p += 2 {
		f(p)
	}
	if x > 1 {
		factors = append(factors, x)
	}

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	for _, p := range factors {
		_, _ = fmt.Fprintln(writer, p)
	}
}
