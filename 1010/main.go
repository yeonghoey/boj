package main

import (
	"bufio"
	"fmt"
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
		n := nextInt()
		m := nextInt()
		a := solve(n, m)
		fmt.Println(a)
		t--
	}

}

func solve(n, m int) int {
	var bs []int
	for i := 0; i < n; i++ {
		bs = append(bs, n-i)
	}

	a := 1
	for i := 0; i < n; i++ {
		a *= m - i
		for j := range bs {
			if a%bs[j] == 0 {
				a /= bs[j]
				bs[j] = 1
			}
		}
	}

	return a
}
