package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	e0, e1 := 0, 1
	for d := 2; d <= N; d++ {
		e0, e1 = e0+e1, e0
	}
	fmt.Println(e0 + e1)
}
