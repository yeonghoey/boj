package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	_, _ = fmt.Scan(&N)

	perm := make([]int, N)
	used := make([]bool, N+1)

	w := bufio.NewWriter(os.Stdout)
	defer func() { _ = w.Flush() }()

	var f func(k int)
	f = func(k int) {
		if k == N {
			s := fmt.Sprint(perm)
			s = strings.Trim(s, "[]")
			_, _ = fmt.Fprintln(w, s)
			return
		}
		for i := 1; i <= N; i++ {
			if !used[i] {
				used[i] = true
				perm[k] = i
				f(k + 1)
				perm[k] = -1
				used[i] = false
			}
		}
	}
	f(0)
}
