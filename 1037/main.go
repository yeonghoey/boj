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

	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = nextInt()
	}

	answer := solve(as)
	fmt.Println(answer)
}

func solve(as []int) int {
	n := len(as)
	asSorted := make([]int, n)
	copy(asSorted, as)
	sort.Ints(asSorted)

	return asSorted[0] * asSorted[n-1]
}
