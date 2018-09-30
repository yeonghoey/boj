package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	a := make([]int, n+1)
	a[1] = 0
	for i := 2; i <= n; i++ {
		v := a[i-1] + 1

		if i%2 == 0 {
			if x := a[i/2] + 1; x < v {
				v = x
			}
		}
		if i%3 == 0 {
			if x := a[i/3] + 1; x < v {
				v = x
			}
		}
		a[i] = v
	}
	fmt.Println(a[n])
}
