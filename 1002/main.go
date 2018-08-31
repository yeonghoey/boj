package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var x1, y1, r1, x2, y2, r2 int
	for i := 0; i < t; i++ {
		fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2)
		distance2 := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
		rdiff2 := (r1 - r2) * (r1 - r2)
		rsum2 := (r1 + r2) * (r1 + r2)

		if distance2 == 0 && rdiff2 == 0 {
			fmt.Println(-1)
		} else if rdiff2 < distance2 && distance2 < rsum2 {
			fmt.Println(2)
		} else if rdiff2 == distance2 || rsum2 == distance2 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
