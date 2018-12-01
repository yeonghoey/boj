package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		a, b := input(reader)
		if a == 0 && b == 0 {
			break
		}
		output(writer, a+b)
	}

	_ = writer.Flush()
}

func input(reader *bufio.Reader) (a, b int) {
	_, _ = fmt.Fscanf(reader, "%d %d\n", &a, &b)
	return
}

func output(writer *bufio.Writer, x int) {
	_, _ = fmt.Fprintln(writer, x)
}
