package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	nextInt := func() int {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		return n
	}

	c := nextInt()
	for c > 0 {
		n := nextInt()
		a := make([]int, n)

		mean := 0.0
		for i := 0; i < n; i++ {
			a[i] = nextInt()
			mean += float64(a[i])
		}
		mean /= float64(n)
		over := 0
		for i := 0; i < n; i++ {
			if float64(a[i]) > mean {
				over++
			}
		}

		answer := float64(over) * 100.0 / float64(n)
		_, _ = fmt.Fprintf(writer, "%.3f%%\n", answer)

		c--
	}
	_ = writer.Flush()
}
