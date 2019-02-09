package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)
	A := make([]int, max(N+1, 4))
	A[1], A[2], A[3] = 0, 1, 1
	for x := 4; x <= N; x++ {
		A[x] = A[x-1]
		if x%2 == 0 && A[x/2] < A[x] {
			A[x] = A[x/2]
		}
		if x%3 == 0 && A[x/3] < A[x] {
			A[x] = A[x/3]
		}
		A[x]++
	}

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()
	_, _ = fmt.Fprintln(w, A[N])
	x := N
	for x > 1 {
		_, _ = fmt.Fprintf(w, "%d ", x)
		n := A[x]
		if x%2 == 0 && A[x/2] == n-1 {
			x /= 2
			continue
		}
		if x%3 == 0 && A[x/3] == n-1 {
			x /= 3
			continue
		}
		x--
	}
	_, _ = fmt.Fprintf(w, "1\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
