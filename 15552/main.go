package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	_, _ = fmt.Fscan(reader, &t)

	for t > 0 {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
		t--
	}
	writer.Flush()
}
