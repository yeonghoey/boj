package main

import "fmt"

func main() {
	var X int
	_, _ = fmt.Scan(&X)
	x := 1
	d := 1
	for x+d <= X {
		x += d
		d++
	}
	diff := X - x
	a, b := 1+diff, d-diff
	if d%2 == 1 {
		a, b = b, a
	}
	fmt.Printf("%d/%d\n", a, b)
}
