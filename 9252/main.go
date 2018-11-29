package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Init
	A := "-" + nextString()
	B := "-" + nextString()

	d := make([][]int, len(A))
	for ai := range d {
		d[ai] = make([]int, len(B))
	}

	// Solve
	for ai := 1; ai < len(A); ai++ {
		for bi := 1; bi < len(B); bi++ {
			a := A[ai]
			b := B[bi]
			if a == b {
				d[ai][bi] = d[ai-1][bi-1] + 1
			} else {
				d[ai][bi] = max(d[ai-1][bi], d[ai][bi-1])
			}
		}
	}

	// Restructure
	ss := make([]byte, 0, max(len(A), len(B)))
	ai := len(A) - 1
	bi := len(B) - 1
	for ai >= 1 && bi >= 1 {
		a := A[ai]
		b := B[bi]

		if a == b {
			ss = append(ss, a)
			ai--
			bi--
		} else {
			if d[ai-1][bi] > d[ai][bi-1] {
				ai--
			} else {
				bi--
			}
		}
	}
	// Reverse
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}

	length := d[len(A)-1][len(B)-1]
	lcs := string(ss)

	fmt.Println(length)
	fmt.Println(lcs)
}

var reader = bufio.NewReader(os.Stdin)

func nextString() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
