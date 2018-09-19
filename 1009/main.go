package main

import (
	"bufio"
	"fmt"
	"os"
)

var cases = [...][]int{
	{10},
	{1},
	{2, 4, 8, 6},
	{3, 9, 7, 1},
	{4, 6},
	{5},
	{6},
	{7, 9, 3, 1},
	{8, 4, 2, 6},
	{9, 1},
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	t := nextInt()
	for t > 0 {
		a := nextInt()
		b := nextInt()
		n := solve(a, b)
		fmt.Println(n)
		t--
	}
}

func solve(a, b int) int {
	aa := a % 10
	cs := cases[aa]
	bb := (b - 1) % len(cs)
	return cs[bb]
}

func nextInt() int {
	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		panic(err)
	}
	return n
}
