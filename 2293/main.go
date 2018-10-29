package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInt := func() int {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		return n
	}

	n := nextInt()
	k := nextInt()
	cases := make([]int, k+1)
	cases[0] = 1
	for i := 0; i < n; i++ {
		coin := nextInt()
		for ki := 1; ki <= k; ki++ {
			if ki-coin >= 0 {
				cases[ki] += cases[ki-coin]
			}
		}
	}
	fmt.Println(cases[k])
}
