package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func input() []int {
	s, _ := reader.ReadString('\n')
	fs := strings.Fields(s)
	ns := make([]int, len(fs))
	for i, f := range fs {
		ns[i], _ = strconv.Atoi(f)
	}
	return ns
}

func output(n int) {
	fmt.Fprintln(writer, n)
}

func main() {
	defer writer.Flush()
	T := input()[0]
	for T > 0 {
		K := input()[0]
		C := input()

		cost := solve(K, C)
		output(cost)
		T--
	}
}

func solve(K int, C []int) int {
	table := [][]int{}
	for i := 0; i < K; i++ {
		table = append(table, make([]int, K))
	}

	for to := 0; to < K; to++ {
		for from := to - 1; from >= 0; from-- {
			x := -1

			for part := from; part < to; part++ {
				left := table[from][part]
				right := table[part+1][to]
				comb := left + right
				if x < 0 || x > comb {
					x = comb
				}
			}

			table[from][to] = x
			for ci := from; ci <= to; ci++ {
				table[from][to] += C[ci]
			}
		}
	}

	return table[0][K-1]
}
