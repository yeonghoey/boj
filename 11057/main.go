package main

import "fmt"

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	A := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 2; i <= N; i++ {
		A1 := make([]int, 10)
		for d, count := range A {
			for d1 := d; d1 < 10; d1++ {
				A1[d1] += count
				A1[d1] %= 10007
			}
		}
		A = A1
	}

	total := 0
	for _, count := range A {
		total += count
		total %= 10007
	}
	fmt.Println(total)
}
