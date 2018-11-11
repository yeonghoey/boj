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

func main() {
	N := nextInt()
	K := nextInt()
	coins := make([]int, N)
	for i := range coins {
		coins[i] = nextInt()
	}

	count, change := 0, K
	for i := N - 1; i >= 0; i-- {
		count += change / coins[i]
		change = change % coins[i]
		if change == 0 {
			break
		}
	}

	fmt.Println(count)
}
