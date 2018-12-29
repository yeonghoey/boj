package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	sign := 1
	x := 0
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
	A := make([]int, N)
	for i := range A {
		A[i] = nextInt()
	}
	sort.Ints(A)
	M := nextInt()
	answer := make([]int, 0)
	for i := 0; i < M; i++ {
		target := nextInt()
		ok := binSearch(A, target)
		if ok {
			answer = append(answer, 1)
		} else {
			answer = append(answer, 0)
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(answer), "[]"))
}

func binSearch(A []int, target int) bool {
	if len(A) < 1 {
		return false
	}

	mid := len(A) / 2

	if A[mid] < target {
		return binSearch(A[mid+1:], target)
	}

	if A[mid] > target {
		return binSearch(A[:mid], target)
	}

	return true
}
