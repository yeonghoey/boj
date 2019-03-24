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
		bs := scanner.Bytes()
		if len(bs) == 1 && bs[0] == '0' {
			break
		}
		lo, hi := 0, len(bs)-1
		for lo < hi && bs[lo] == bs[hi] {
			lo++
			hi--
		}
		answer := "no"
		if bs[lo] == bs[hi] {
			answer = "yes"
		}
		_, _ = fmt.Fprintln(writer, answer)
	}
	_ = writer.Flush()
}
