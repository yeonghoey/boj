package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	t := nextInt()
	for t > 0 {
		x := nextInt()
		y := nextInt()
		count := solve(x, y)
		fmt.Println(count)
		t--
	}
}

func solve(x, y int) int {
	n := y - x
	half1 := int(math.Floor(root(n / 2)))
	m := n - (half1 * (half1 + 1) / 2)
	half2 := int(math.Ceil(root(m)))
	return half1 + half2
}

func root(n int) float64 {
	return (-1 + math.Sqrt(1+float64(8*n))) / 2
}
