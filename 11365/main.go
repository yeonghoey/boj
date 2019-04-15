package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "END" {
			break
		}
		n := len(line)
		rs := make([]rune, n)
		for i, r := range line {
			rs[n-i-1] = r
		}
		_, _ = fmt.Fprintln(writer, string(rs))
	}
	_ = writer.Flush()
}
