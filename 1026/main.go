package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInt := func() int {
		var n int
		_, err := fmt.Fscan(reader, &n)
		if err != nil {
			panic(err)
		}
		return n
	}

	n := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		b[i] = nextInt()
	}

	s := solve(a, b)
	fmt.Println(s)
}

func solve(a, b []int) int {
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	s := 0
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}
