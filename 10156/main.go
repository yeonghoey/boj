package main

import "fmt"

func main() {
	var K, N, M int
	fmt.Scan(&K, &N, &M)
	x := K*N - M
	if x < 0 {
		x = 0
	}
	fmt.Println(x)
}
