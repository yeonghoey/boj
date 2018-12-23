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
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

type elem struct {
	value int
	index int
}

type withIndex []elem

func (a withIndex) Len() int {
	return len(a)
}

func (a withIndex) Less(i, j int) bool {
	if a[i].value != a[j].value {
		return a[i].value < a[j].value
	}
	return a[i].index < a[j].index
}

func (a withIndex) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	N := nextInt()
	A := make([]elem, N)
	for i := range A {
		A[i].value = nextInt()
		A[i].index = i
	}

	sort.Sort(withIndex(A))

	pMap := make(map[int]int)
	for i, e := range A {
		pMap[e.index] = i
	}

	P := make([]int, N)
	for i := range P {
		P[i] = pMap[i]
	}

	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(writer, "%d", P[0])
	for i := 1; i < N; i++ {
		fmt.Fprintf(writer, " %d", P[i])
	}
	_ = writer.Flush()
}
