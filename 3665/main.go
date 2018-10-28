package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	a, b int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nextInts := func() []int {
		scanner.Scan()
		line := scanner.Text()
		fields := strings.Fields(line)
		ns := make([]int, len(fields))
		for i, f := range fields {
			ns[i], _ = strconv.Atoi(f)
		}
		return ns
	}

	writer := bufio.NewWriter(os.Stdout)
	T := nextInts()[0]

	for T > 0 {
		n := nextInts()[0]
		t := nextInts()
		m := nextInts()[0]
		pairs := make([]pair, m)
		for i := range pairs {
			ab := nextInts()
			pairs[i].a = ab[0]
			pairs[i].b = ab[1]
		}

		answer := solve(n, t, m, pairs)
		_, _ = fmt.Fprintln(writer, answer)
		T--
	}

	_ = writer.Flush()
}

func solve(n int, t []int, m int, pairs []pair) string {
	g := make(map[int]map[int]bool)
	in := make(map[int]int)
	for _, node := range t {
		g[node] = make(map[int]bool)
		in[node] = 0
	}

	for i := 0; i < n; i++ {
		head := t[i]
		for j := i + 1; j < n; j++ {
			tail := t[j]
			g[tail][head] = true
			in[head]++
		}
	}

	for _, p := range pairs {
		t, h := p.a, p.b
		if g[p.b][p.a] {
			t, h = p.b, p.a
		}
		delete(g[t], h)
		in[h]--
		g[h][t] = true
		in[t]++
	}

	sources := make([]int, 0, n)
	for head, count := range in {
		if count == 0 {
			sources = append(sources, head)
		}
	}

	ranks := make([]int, 0, n)
	explored := make(map[int]bool)
	visited := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(node int) bool {
		explored[node] = true
		visited[node] = true
		for head := range g[node] {
			// cycle
			if visited[head] {
				return false
			}
			if !explored[head] {
				if ok := dfs(head); !ok {
					return false
				}
			}
		}
		visited[node] = false
		ranks = append(ranks, node)
		return true
	}

	answer := "IMPOSSIBLE"
	if len(sources) == 1 {
		start := sources[0]
		ok := dfs(start)
		if ok && len(ranks) == n {
			answer = format(ranks)
		}
	}

	return answer
}

func format(ns []int) string {
	ss := make([]string, len(ns))
	for i, n := range ns {
		ss[i] = strconv.Itoa(n)
	}

	return strings.Join(ss, " ")
}
