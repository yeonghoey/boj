package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const maxN = 10000

func main() {
	writer := bufio.NewWriter(os.Stdout)

	a := make([]bool, maxN+1)
	for i := 1; i <= maxN; i++ {
		if !a[i] {
			_, _ = fmt.Fprintln(writer, i)
		}
		x := d(i)
		if x <= maxN {
			a[x] = true
		}
	}

	_ = writer.Flush()
}

func d(n int) int {
	dnum := n
	for _, d := range strconv.Itoa(n) {
		dnum += int(d - '0')
	}
	return dnum
}
