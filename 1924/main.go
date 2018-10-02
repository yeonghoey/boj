package main

import (
	"fmt"
)

func main() {
	days := [13]int{-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	weekdays := [7]string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	var m, d int
	_, _ = fmt.Scan(&m, &d)

	count := 0
	for i := 1; i <= m; i++ {
		dd := days[i]
		if i == m {
			dd = d
		}
		for j := 1; j <= dd; j++ {
			count++
		}
	}

	fmt.Println(weekdays[count%7])
}
