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

	m := next(n)
	cycle := 1
	for n != m {
		m = next(m)
		cycle++
	}

	fmt.Println(cycle)
}

func next(m int) int {
	x := m

	a, b := x/10, x%10
	c := a + b

	return b*10 + c%10
}
