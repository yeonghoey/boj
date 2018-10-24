package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInt := func() int {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		return n
	}

	N := nextInt()
	A := make([]struct{ x, y int }, N)
	for i := 0; i < N; i++ {
		A[i].x = nextInt()
		A[i].y = nextInt()
	}

	ranks := make([]string, N)
	for i, ai := range A {
		count := 1
		for j, aj := range A {
			if i == j {
				continue
			}
			if ai.x < aj.x && ai.y < aj.y {
				count++
			}
		}
		ranks[i] = strconv.Itoa(count)
	}

	fmt.Println(strings.Join(ranks, " "))
}
