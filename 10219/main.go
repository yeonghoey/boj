package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	var T int
	_, _ = fmt.Fscanf(reader, "%d\n", &T)

	for T > 0 {
		var H, W int
		_, _ = fmt.Fscanf(reader, "%d %d\n", &H, &W)
		for H > 0 {
			s, _ := reader.ReadString('\n')
			s = strings.TrimSpace(s)
			s = reverse(s)
			_, _ = fmt.Fprintln(writer, s)
			H--
		}
		T--
	}
	_ = writer.Flush()
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
