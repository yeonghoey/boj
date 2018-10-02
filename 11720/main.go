package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	_, _ = fmt.Scan(&n, &s)
	ans := 0
	for i := 0; i < n; i++ {
		ans += int(s[i] - '0')
	}
	fmt.Println(ans)
}
