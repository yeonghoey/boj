package main

import (
	"fmt"
)

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	a, b := A, B
	if A > B {
		a, b = b, a
	}

	round := 1
	for !(a%2 == 1 && b-a == 1) {
		a = (a + 1) / 2
		b = (b + 1) / 2
		round++
	}
	fmt.Println(round)
}
