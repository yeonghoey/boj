package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	line := strings.Repeat("*", n)
	for i := 0; i < n; i++ {
		fmt.Println(line)
	}
}
