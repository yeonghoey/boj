package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxN = 1001
const maxK = 100001

var costs [maxN]int
var totalCosts [maxN]int
var depCounts [maxN]int
var deps [maxN][maxN]int

var stack [maxK]int
var visited [maxN]bool
var order [maxN]int
var ordered [maxN]bool

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var T int
	mustScan(&T)
	for i := 0; i < T; i++ {
		t := process()
		fmt.Println(t)
	}
}

func process() int {
	var N, K, W int

	mustScan(&N, &K)

	for i := 1; i <= N; i++ {
		mustScan(&costs[i])
		totalCosts[i] = costs[i]
		depCounts[i] = 0
		visited[i] = false
		ordered[i] = false
	}

	for i := 0; i < K; i++ {
		var pre, post int
		mustScan(&pre, &post)
		idx := depCounts[post]
		depCounts[post]++
		deps[post][idx] = pre
	}

	mustScan(&W)

	stackN, orderN := 1, 0
	stack[0] = W
	for stackN > 0 {
		current := stack[stackN-1]

		if visited[current] {
			if !ordered[current] {
				ordered[current] = true
				order[orderN] = current
				orderN++
			}
			stackN--
			continue
		}

		for i := 0; i < depCounts[current]; i++ {
			pre := deps[current][i]
			if !visited[pre] {
				stack[stackN] = pre
				stackN++
			}
		}

		visited[current] = true
	}

	for i := 0; i < orderN; i++ {
		target := order[i]
		maxDepCost := 0
		for j := 0; j < depCounts[target]; j++ {
			pre := deps[target][j]
			depCost := totalCosts[pre]
			if depCost > maxDepCost {
				maxDepCost = depCost
			}
		}
		totalCosts[target] = costs[target] + maxDepCost
	}

	return totalCosts[W]
}

func mustScan(a ...interface{}) {
	_, err := fmt.Fscan(reader, a...)
	if err != nil {
		panic(err)
	}
}
