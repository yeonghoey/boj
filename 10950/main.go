package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func nextInt() int {
	var n int
	_, _ = fmt.Fscanf(reader, "%d ", &n)
	return n
}

func main() {
	T := nextInt()
	for T > 0 {
		A := nextInt()
		B := nextInt()
		_, _ = fmt.Fprintln(writer, A+B)
		T--
	}
	_ = writer.Flush()
}
