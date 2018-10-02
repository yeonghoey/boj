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
		for j := n - i + 1; j >= 1; j-- {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}
	writer.Flush()
}
