package main

import "fmt"

func main() {
	var x, y, w, h int
	_, err := fmt.Scan(&x, &y, &w, &h)

	if err != nil {
		panic(err)
	}

	ds := [...]int{x, y, w - x, h - y}
	min := ds[0]
	for _, d := range ds {
		if d < min {
			min = d
		}
	}

	fmt.Println(min)
}
