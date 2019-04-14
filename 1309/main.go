package main

import "fmt"

const m = 9901

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	a, b, c := 1, 1, 1
	for i := 2; i <= N; i++ {
		a, b, c = (b+c)%m, (a+c)%m, (a+b+c)%m
	}
	fmt.Println((a + b + c) % m)
}
