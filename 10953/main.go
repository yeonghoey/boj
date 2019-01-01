package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	var T int
	_, _ = fmt.Fscanf(reader, "%d\n", &T)
	for T > 0 {
		var A, B int
		_, _ = fmt.Fscanf(reader, "%d,%d\n", &A, &B)
		_, _ = fmt.Fprintln(writer, A+B)
		T--
	}
}
