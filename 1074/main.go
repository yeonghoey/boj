package main

import "fmt"

var exp2 []int

func main() {
	N, r, c := input()
	cacheExp2(N * 2)
	answer := search(N, r, c, 0)
	fmt.Println(answer)
}

func input() (N int, r int, c int) {
	_, _ = fmt.Scan(&N, &r, &c)
	return
}

func cacheExp2(N int) {
	exp2 = make([]int, N+1)
	exp2[0] = 1
	for i := 1; i <= N; i++ {
		exp2[i] = exp2[i-1] * 2
	}
}

func search(n, r, c, k int) int {
	if n == 0 {
		return k
	}

	h := exp2[n-1]
	s := h * h

	// 1st quadrant
	if r < h && c < h {
		return search(n-1, r, c, k)
	}

	// 2nd quadrant
	if r < h && c >= h {
		return search(n-1, r, c-h, k+s)
	}

	// 3rd quadrant
	if r >= h && c < h {
		return search(n-1, r-h, c, k+s*2)
	}

	// 4th quadrant
	if r >= h && c >= h {
		return search(n-1, r-h, c-h, k+s*3)
	}

	// Unexpected
	return -1
}
