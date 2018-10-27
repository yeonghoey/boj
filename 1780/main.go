package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		ns := make([]int, len(fields))
		for i, f := range fields {
			ns[i], _ = strconv.Atoi(f)
		}
		return ns
	}

	N := nextInts()[0]
	A := make([][]int, N)
	for i := range A {
		row := nextInts()
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			// (-1, 0, 1) -> (0, 1, 2)
			A[i][j] = row[j] + 1
		}
	}

	answer := count(A, 0, N, 0, N)
	fmt.Println(answer[0])
	fmt.Println(answer[1])
	fmt.Println(answer[2])
}

func count(A [][]int, y0, y1, x0, x1 int) [3]int {
	// Test whether the current paper is singleton
	c := [3]int{0, 0, 0}
	for yi := y0; yi < y1; yi++ {
		for xi := x0; xi < x1; xi++ {
			a := A[yi][xi]
			c[a] = 1
			if c[0]+c[1]+c[2] > 1 {
				break
			}
		}
		if c[0]+c[1]+c[2] > 1 {
			break
		}
	}

	if c[0]+c[1]+c[2] == 1 {
		return c
	}

	// Collect counts of subproblems
	s := [3]int{0, 0, 0}
	yd, xd := (y1-y0)/3, (x1-x0)/3
	for yi := y0; yi < y1; yi += yd {
		for xi := x0; xi < x1; xi += xd {
			ss := count(A, yi, yi+yd, xi, xi+xd)
			s[0] += ss[0]
			s[1] += ss[1]
			s[2] += ss[2]
		}
	}
	return s
}
