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
func nextInt() int {
	scanner.Scan()
	sign, x := 1, 0
	for _, b := range scanner.Bytes() {
		if b == '-' {
			sign = -1
			continue
		}
		x *= 10
		x += int(b - '0')
	}
	return sign * x
}

func main() {
	E := nextInt()
	S := nextInt()
	M := nextInt()
	var answer int
	for y := 0; y < 7980; y++ {
		if true &&
			((y%15)+1 == E) &&
			((y%28)+1 == S) &&
			((y%19)+1 == M) {
			answer = y + 1
			break
		}
	}
	fmt.Println(answer)
}
