package main

import (
	"fmt"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	prev := '|'
	height := 0
	for _, x := range s {
		if x != prev {
			height += 10
		} else {
			height += 5
		}
		prev = x
	}
	fmt.Println(height)
}
