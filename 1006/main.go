package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	top = 0
	bot = 1
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input = scanner.Text()
	t, _ := strconv.Atoi(input)

	for t > 0 {
		scanner.Scan()
		input = scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		n, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])

		var a [][]int
		for i := 0; i < 2; i++ {
			a = append(a, make([]int, n))
			scanner.Scan()
			input = scanner.Text()
			parts := strings.SplitN(input, " ", n)
			for j := 0; j < n; j++ {
				x, _ := strconv.Atoi(parts[j])
				a[i][j] = x
			}
		}
		answer := solve(a, n, w)
		fmt.Println(answer)

		t--
	}
}

func solve(a [][]int, n, w int) int {
	aa := [][]int{make([]int, n+1), make([]int, n+1)}
	copy(aa[top][1:], a[top])
	copy(aa[bot][1:], a[bot])

	answer := 2 * n

	aa[top][0] = w
	aa[top][n] = a[top][n-1]
	aa[bot][0] = w
	aa[bot][n] = a[bot][n-1]
	answer = min(answer, solveCase(aa, n, w))
	aa[top][0] = a[top][n-1]
	aa[top][n] = w
	aa[bot][0] = w
	aa[bot][n] = a[bot][n-1]
	answer = min(answer, solveCase(aa, n, w))
	aa[top][0] = a[top][n-1]
	aa[top][n] = w
	aa[bot][0] = a[bot][n-1]
	aa[bot][n] = w
	answer = min(answer, solveCase(aa, n, w))
	aa[top][0] = w
	aa[top][n] = a[top][n-1]
	aa[bot][0] = a[bot][n-1]
	aa[bot][n] = w
	answer = min(answer, solveCase(aa, n, w))
	return answer
}

func solveCase(aa [][]int, n, w int) int {
	max := len(aa) * len(aa[0])
	topFree := 2
	botFree := 2
	allFree := 2
	notFree := max

	if aa[top][0]+aa[bot][1] <= w {
		notFree = 2
	}

	for i := 1; i <= n; i++ {
		sumTop := aa[top][i-1] + aa[top][i]
		sumBot := aa[bot][i-1] + aa[bot][i]
		sumCol := aa[top][i] + aa[bot][i]

		topJoin := max
		if sumTop <= w {
			topJoin = min(allFree, topFree) + 1
		}
		botJoin := max
		if sumBot <= w {
			botJoin = min(allFree, botFree) + 1
		}
		allJoin := max
		if sumTop <= w && sumBot <= w {
			allJoin = min(allJoin, allFree)
		}
		if sumCol <= w {
			colJoin := min(topFree, botFree, allFree, notFree) + 1
			allJoin = min(allJoin, colJoin)
		}
		notJoin := min(topFree, botFree, notFree) + 2

		topFree = min(botJoin, notJoin)
		botFree = min(topJoin, notJoin)
		allFree = notJoin
		notFree = allJoin
	}

	best := min(topFree, botFree, allFree, notFree) - 2
	return best
}

func min(a int, as ...int) int {
	m := a
	for _, x := range as {
		if x < m {
			m = x
		}
	}
	return m
}
