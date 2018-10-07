package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	maxNumber = 10000
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	nextInt := func() int {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		return n
	}

	a := make([]int, maxNumber+1)
	n := nextInt()
	for i := 0; i < n; i++ {
		a[nextInt()]++
	}

	for i := 1; i <= maxNumber; i++ {
		for j := 0; j < a[i]; j++ {
			_, _ = fmt.Fprintln(writer, i)
		}
	}
	_ = writer.Flush()
}
