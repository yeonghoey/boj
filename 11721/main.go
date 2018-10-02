package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	var s string
	_, _ = fmt.Scan(&s)

	for i := 0; i <= len(s)/10; i++ {
		b := i * 10
		e := (i + 1) * 10
		if e > len(s) {
			e = len(s)
		}
		fmt.Fprintln(writer, s[b:e])
	}
	writer.Flush()
}
