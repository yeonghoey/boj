package main

import (
	"bufio"
	"fmt"
	"os"
)

type color struct {
	red   int
	green int
	blue  int
}

func min(n int, ns ...int) int {
	m := n
	for _, nn := range ns {
		if nn < m {
			m = nn
		}
	}
	return m
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInt := func() int {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		return n
	}

	n := nextInt()
	costs := make([]color, n)
	for i := 0; i < n; i++ {
		costs[i].red = nextInt()
		costs[i].green = nextInt()
		costs[i].blue = nextInt()

		if i > 0 {
			prev := costs[i-1]
			costs[i].red += min(prev.green, prev.blue)
			costs[i].green += min(prev.red, prev.blue)
			costs[i].blue += min(prev.red, prev.green)
		}

	}

	answer := min(costs[n-1].red, costs[n-1].green, costs[n-1].blue)
	fmt.Println(answer)
}
