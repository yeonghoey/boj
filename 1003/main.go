package main

import (
	"fmt"
)

func fibc(N int) (int, int) {
	a0, a1 := 1, 0
	b0, b1 := 0, 1

	if N == 0 {
		return a0, a1
	} else if N == 1 {
		return b0, b1
	}

	var c0, c1 int
	for i := 2; i <= N; i++ {
		c0, c1 = a0+b0, a1+b1
		a0, a1 = b0, b1
		b0, b1 = c0, c1
	}

	return c0, c1
}

func main() {
	var T, N int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		c0, c1 := fibc(N)
		fmt.Println(c0, c1)
	}
}
