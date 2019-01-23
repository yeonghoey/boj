package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	x := N
	factors := make([]int, 0)
	isComposite := make([]bool, N+1)
	for p := 2; p <= N; p++ {
		if !isComposite[p] {
			for x%p == 0 {
				factors = append(factors, p)
				x /= p
			}

			for y := p + p; y <= x; y += p {
				isComposite[y] = true
			}
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	for _, p := range factors {
		_, _ = fmt.Fprintln(writer, p)
	}
}
