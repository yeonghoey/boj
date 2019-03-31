package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func next() (bs []byte, ok bool) {
	ok = scanner.Scan()
	bs = scanner.Bytes()
	return
}

func main() {
	var lines [][]byte
	for {
		bs, ok := next()
		if !ok {
			break
		}

		line := make([]byte, 15)
		copy(line, bs)
		lines = append(lines, line)
	}

	var buf []byte
	for i := 0; i < 15; i++ {
		for j := range lines {
			ch := lines[j][i]
			if ch != 0 {
				buf = append(buf, ch)
			}
		}
	}
	fmt.Println(string(buf))
}
