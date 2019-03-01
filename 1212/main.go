package main

import (
	"bufio"
	"os"
	"strings"
)

var o2b = [...]string{"000", "001", "010", "011", "100", "101", "110", "111"}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
}

func next() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	first := true
	for {
		d := next()
		d -= '0'
		if !(0 <= d && d < 8) {
			break
		}
		s := o2b[d]
		if first {
			first = false
			s = strings.TrimLeft(s, "0")
		}
		_, _ = w.WriteString(s)
	}
	if w.Buffered() == 0 {
		_ = w.WriteByte('0')
	}
}
