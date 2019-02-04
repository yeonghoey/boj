package main

import (
	"bufio"
	"fmt"
	"os"
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
	T := nextInt()
	for T > 0 {
		handleCase()
		T--
	}
}

func handleCase() {
	n := nextInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}

	order := []int{}
	visited := make([]bool, n+1)

	var dfs func(int)
	dfs = func(x int) {
		if !visited[x] {
			visited[x] = true
			dfs(a[x])
			order = append(order, x)
		}
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	ra := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		x := a[i]
		ra[x] = append(ra[x], i)
	}

	assigned := make([]bool, n+1)
	counter := make(map[int]int)

	var assign func(int, int)
	assign = func(x, root int) {
		if !assigned[x] {
			assigned[x] = true
			counter[root]++
			for _, y := range ra[x] {
				assign(y, root)
			}
		}
	}

	for i := len(order) - 1; i >= 0; i-- {
		x := order[i]
		if !assigned[x] {
			assign(x, x)
		}
	}

	answer := 0
	for root, count := range counter {
		if count == 1 && root != a[root] {
			answer += count
		}
	}

	fmt.Println(answer)
}
