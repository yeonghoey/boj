package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var x1, y1, r1, x2, y2, r2 int
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %d %d %d %d", &x1, &y1, &r1, &x2, &y2, &r2)
		diff2 := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
		rsub2 := (r1 - r2) * (r1 - r2)
		rsum2 := (r1 + r2) * (r1 + r2)

		if diff2 == 0 && rsub2 == 0 {
			fmt.Println(-1)
		} else if rsub2 < diff2 && diff2 < rsum2 {
			fmt.Println(2)
		} else if rsub2 == diff2 || rsum2 == diff2 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
