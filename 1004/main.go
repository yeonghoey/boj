package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var x1, y1, x2, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)
		var n int
		fmt.Scan(&n)
		count := 0
		for j := 0; j < n; j++ {
			var cx, cy, r int
			fmt.Scan(&cx, &cy, &r)
			in1 := isEncircled(x1, y1, cx, cy, r)
			in2 := isEncircled(x2, y2, cx, cy, r)
			if (in1 && !in2) || (!in1 && in2) {
				count += 1
			}
		}
		fmt.Println(count)
	}
}

func isEncircled(x, y, cx, cy, r int) bool {
	distance2 := (x-cx)*(x-cx) + (y-cy)*(y-cy)
	return distance2 < r*r
}
