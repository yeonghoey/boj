package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	sign, x := 1, 0
	for _, b := range scanner.Bytes() {
		if b == '-' {
			sign = -1
			continue
		}
		x *= 10
		x += int(b - '0')
	}
	return sign * x
}

func main() {
	N := nextInt()
	S := nextInt()
	A := make([]int, N)
	for i := range A {
		A[i] = nextInt()
	}
	sums1 := computeSums(A[:N/2])
	sums2 := computeSums(A[N/2:])

	count := 0
	// Can make S with only the first half
	for _, s := range sums1 {
		if s == S {
			count++
		}
	}
	// Can make S with only the second half
	for _, s := range sums2 {
		if s == S {
			count++
		}
	}

	// Can make S with some of the first half and some of the second half
	sort.Ints(sums1)
	sort.Ints(sums2)
	i1 := 0
	i2 := len(sums2) - 1
	for i1 < len(sums1) && i2 >= 0 {
		s1 := sums1[i1]
		s2 := sums2[i2]
		s := s1 + s2
		if s == S {
			var d1, d2 int
			for d1 = i1; d1 < len(sums1) && sums1[d1] == s1; d1++ {
			}
			for d2 = i2; d2 >= 0 && sums2[d2] == s2; d2-- {
			}
			count += (d1 - i1) * (i2 - d2)
			i1, i2 = d1, d2
		} else if s < S {
			i1++
		} else {
			i2--
		}
	}
	fmt.Println(count)
}

func computeSums(nums []int) []int {
	sums := []int{}
	var f func(int, int, int)
	f = func(k, s, picked int) {
		if k == len(nums) {
			if picked > 0 {
				sums = append(sums, s)
			}
			return
		}
		f(k+1, s, picked)
		f(k+1, s+nums[k], picked+1)
	}
	f(0, 0, 0)
	return sums
}
