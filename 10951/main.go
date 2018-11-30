package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		a, b, eof := input(reader)
		if eof {
			break
		}
		output(writer, a+b)
	}

	_ = writer.Flush()
}

func input(reader *bufio.Reader) (a, b int, eof bool) {
	_, err := fmt.Fscanf(reader, "%d %d\n", &a, &b)
	if err == io.EOF {
		eof = true
	}
	return
}

func output(writer *bufio.Writer, x int) {
	_, _ = fmt.Fprintln(writer, x)
}
