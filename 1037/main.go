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

	lo := asSorted[0]
	hi := asSorted[n-1]

	answer := 1
	for i, a := range asSorted {
		if a == 1 {
			continue
		}

		answer *= a

		for j := i + 1; j < n; j++ {
			if asSorted[j]%a == 0 {
				asSorted[j] /= a
			}
		}
	}

	if answer == hi {
		answer *= lo
	}

	return answer
}

func solveEasy(as []int) int {
	n := len(as)
	asSorted := make([]int, n)
	copy(asSorted, as)
	sort.Ints(asSorted)

	lo := asSorted[0]
	hi := asSorted[n-1]

	return lo * hi
}
