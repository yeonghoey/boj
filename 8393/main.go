package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println(n * (n + 1) / 2)
}
