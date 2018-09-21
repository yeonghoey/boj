package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	n := nextInt()
	answer := solve(n)
	fmt.Println(answer)
}

func solve(n int) int {
	if n < 100 {
		return n
	}

	x := n
	if n == 1000 {
		x = 999
	}

	h := x / 100
	count := 0
	for i := 1; i <= h; i++ {
		for j := -i / 2; j <= (9-i)/2; j++ {
			a := i * 100
			b := (i + j) * 10
			c := (i + j + j)
			if a+b+c <= x {
				count++
			}
		}
	}
	return count + 99
}

func nextInt() int {
	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		panic(err)
	}
	return n
}
