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
	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
