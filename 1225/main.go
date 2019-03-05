package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)
}

func nextDigit() int {
	ok := scanner.Scan()
	if !ok {
		return -1
	}
	b := scanner.Bytes()[0]
	if !('0' <= b && b <= '9') {
		return -1
	}
	return int(b - '0')
}

func main() {
	a := 0
	for {
		if d := nextDigit(); d >= 0 {
			a += d
		} else {
			break
		}
	}
	b := 0
	for {
		if d := nextDigit(); d >= 0 {
			b += d
		} else {
			break
		}
	}
	fmt.Println(a * b)
}
