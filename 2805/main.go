package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	N := nextInt()
	M := nextInt()

	woods := make([]int, N)
	for i := range woods {
		woods[i] = nextInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(woods)))

	cumsum := make([]int, N)
	sum := 0
	for i, wood := range woods {
		sum += wood
		cumsum[i] = sum
	}

	check := func(h int) bool {
		idx := atLeast(woods, h, 0, N)
		m := cumsum[idx] - h*(idx+1)
		return m >= M
	}

	var answer int
	var find func(int, int)

	find = func(begin, end int) {
		if begin >= end {
			return
		}

		mid := (begin + end) / 2
		if check(mid) {
			if mid > answer {
				answer = mid
			}
			find(mid+1, end)
		} else {
			find(begin, mid)
		}
	}

	tallest := woods[0]
	find(0, tallest+1)

	fmt.Println(answer)
}

func atLeast(reversed []int, h, begin, end int) int {
	if begin == end {
		return end - 1
	}

	mid := (begin + end) / 2
	if reversed[mid] < h {
		return atLeast(reversed, h, begin, mid)
	} else if reversed[mid] >= h {
		return atLeast(reversed, h, mid+1, end)
	}

	panic("Unexpected")
}
