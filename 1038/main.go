package main

import "fmt"

func main() {
	A := decreasingNumbers()
	var N int
	_, _ = fmt.Scan(&N)
	answer := -1
	if N < len(A) {
		answer = A[N]
	}
	fmt.Println(answer)
}

func decreasingNumbers() []int {
	var walk func(int, int)

	numbers := make([]int, 0)
	digits := make([]int, 10)
	walk = func(di, dn int) {
		if di == dn {
			number, e := 0, 1
			for i := dn - 1; i >= 0; i-- {
				number += digits[i] * e
				e *= 10
			}
			numbers = append(numbers, number)
			return
		}

		var last int
		if di > 0 {
			last = digits[di-1]
		} else {
			last = 10
		}

		for d := 0; d < last; d++ {
			digits[di] = d
			walk(di+1, dn)
		}
	}

	for dn := 1; dn <= 10; dn++ {
		walk(0, dn)
	}

	return numbers
}
