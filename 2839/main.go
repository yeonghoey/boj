package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	found := false
	for q := n / 5; q >= 0; q-- {
		r := n - q*5
		if r%3 == 0 {
			found = true
			fmt.Println(q + r/3)
			break
		}
	}

	if !found {
		fmt.Println(-1)
	}
}
