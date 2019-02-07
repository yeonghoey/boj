package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(nil, 100010)
}

func read() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func main() {
	bytes := read()
	left, count := 0, 0
	var prev byte
	for _, b := range bytes {
		switch b {
		case '(':
			left++
		case ')':
			left--
			switch prev {
			case '(':
				count += left
			case ')':
				count++
			}
		}
		prev = b
	}
	fmt.Println(count)
}
