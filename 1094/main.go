package main

import (
	"bufio"
	"fmt"
	"os"
)

var sticks = [...]int{64, 32, 16, 8, 4, 2, 1}
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	x := nextInt()
	count := solve(x)
	fmt.Println(count)
}

func solve(x int) int {
	count, remain := 0, x
	for _, s := range sticks {
		if remain >= s {
			count++
			remain %= s
		}
		if remain%s == 0 {
			return count
		}
	}

	panic("Unexpected")
}

func nextInt() int {
	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		panic(err)
	}
	return n
}
