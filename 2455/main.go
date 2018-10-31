package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func nextInt() int {
	var n int
	_, _ = fmt.Fscan(reader, &n)
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	current, most := 0, 0
	for i := 0; i < 4; i++ {
		current -= nextInt()
		current += nextInt()
		most = max(most, current)
	}
	fmt.Println(most)
}
