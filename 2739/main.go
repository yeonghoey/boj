package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= 9; i++ {
		fmt.Fprintf(writer, "%d * %d = %d\n", n, i, n*i)
	}
	writer.Flush()
}
