package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	_, _ = fmt.Fscan(reader, &n)

	m := -1
	var scores []int
	for i := 0; i < n; i++ {
		var x int
		_, _ = fmt.Fscan(reader, &x)
		scores = append(scores, x)

		if x > m {
			m = x
		}

	}

	mean := 0.0
	for _, x := range scores {
		mean += (100.0 * float64(x)) / float64(m)
	}
	mean /= float64(n)

	fmt.Printf("%.2f\n", mean)
}
