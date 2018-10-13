package main

import (
	"bufio"
	"fmt"
	"os"
)

var ladder [30 + 1][10 + 1]bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInt := func() int {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		return n
	}

	n := nextInt()
	m := nextInt()
	h := nextInt()

	for i := 0; i < m; i++ {
		a := nextInt()
		b := nextInt()
		ladder[a][b] = true
	}

	answer := solve(n, m, h)
	fmt.Println(answer)
}

func solve(n, m, h int) int {
	answer := -1
	for k := m % 2; k <= 3; k += 2 {
		if track(n, h, k) {
			answer = k
			break
		}
	}

	return answer
}

func track(n, h, k int) bool {
	var try func(int, int, int) bool
	try = func(a1, bnext, added int) bool {
		if added >= k {
			return run(n, h)
		}
		for a := a1; a <= h; a++ {
			b1 := 1
			if a == a1 {
				b1 = bnext
			}
			for b := b1; b < n; b++ {
				if ladder[a][b-1] || ladder[a][b+0] || ladder[a][b+1] {
					continue
				}
				ladder[a][b] = true
				ok := try(a, b+2, added+1)
				ladder[a][b] = false
				if ok {
					return true
				}
			}
		}
		return false
	}
	return try(1, 1, 0)
}

func run(n, h int) bool {
	ok := true
	for b := 1; b <= n; b++ {
		pos := b
		for a := 1; a <= h; a++ {
			if ladder[a][pos] {
				pos++
			} else if ladder[a][pos-1] {
				pos--
			}
		}
		if pos != b {
			ok = false
			break
		}
	}
	return ok
}
