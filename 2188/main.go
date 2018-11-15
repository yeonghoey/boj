package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func nextInts() []int {
	s, _ := reader.ReadString('\n')
	fs := strings.Fields(s)
	ns := make([]int, len(fs))
	for i, f := range fs {
		ns[i], _ = strconv.Atoi(f)
	}
	return ns
}

func main() {
	line1 := nextInts()

	N := line1[0]
	cows := make([][]int, N)
	for i := range cows {
		ns := nextInts()
		cows[i] = ns[1:]
	}

	M := line1[1]
	sheds := make([]int, M+1)
	for i := range sheds {
		sheds[i] = -1
	}

	var pick func(int, []bool) bool
	pick = func(ci int, visited []bool) bool {
		if visited[ci] {
			return false
		}
		visited[ci] = true

		candidates := cows[ci]
		for _, shid := range candidates {
			prev := sheds[shid]
			if prev == -1 || pick(prev, visited) {
				sheds[shid] = ci
				return true
			}
		}
		return false
	}

	count := 0
	for ci := range cows {
		visited := make([]bool, N)
		if pick(ci, visited) {
			count++
		}
	}
	fmt.Println(count)
}
