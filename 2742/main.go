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
	for i := n; i >= 1; i-- {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
