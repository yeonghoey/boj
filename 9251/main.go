package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	A := "-" + nextString()
	B := "-" + nextString()

	d := make([][]int, len(A))
	for i := range d {
		d[i] = make([]int, len(B))
	}

	for i := 1; i < len(A); i++ {
		a := A[i]
		for j := 1; j < len(B); j++ {
			b := B[j]

			if a == b {
				d[i][j] = d[i-1][j-1] + 1
			} else {
				d[i][j] = max(d[i-1][j], d[i][j-1])
			}
		}
	}

	answer := d[len(A)-1][len(B)-1]
	fmt.Println(answer)
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
