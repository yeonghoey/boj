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

func max(n int, ns ...int) int {
	x := n
	for _, n1 := range ns {
		if n1 > x {
			x = n1
		}
	}
	return x
}

func main() {
	n := nextInt()
	a := make([]int, n+3)
	prev, wine := 0, 0
	for i := 3; i <= n+2; i++ {
		wine = nextInt()
		a[i] = max(a[i-1], a[i-2]+wine, a[i-3]+prev+wine)
		prev = wine
	}
	fmt.Println(a[n+2])
}
