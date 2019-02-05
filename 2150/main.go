package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	V := nextInt()
	E := nextInt()

	Gn := make([][]int, V+1)
	Gr := make([][]int, V+1)
	for i := 0; i < E; i++ {
		tail := nextInt()
		head := nextInt()
		Gn[tail] = append(Gn[tail], head)
		Gr[head] = append(Gr[head], tail)
	}

	visited := make([]bool, V+1)
	order := make([]int, 0)

	var visit func(int)
	visit = func(x int) {
		if !visited[x] {
			visited[x] = true
			for _, h := range Gn[x] {
				visit(h)
			}
			order = append(order, x)
		}
	}

	for i := 1; i <= V; i++ {
		if !visited[i] {
			visit(i)
		}
	}

	assigned := make([]bool, V+1)
	components := make([][]int, V+1)

	var assign func(int, int)
	assign = func(x, root int) {
		if !assigned[x] {
			assigned[x] = true
			components[root] = append(components[root], x)
			for _, in := range Gr[x] {
				assign(in, root)
			}
		}
	}

	for i := len(order) - 1; i >= 0; i-- {
		x := order[i]
		if !assigned[x] {
			assign(x, x)
		}
	}

	table := make([][]int, 0)
	for i := 1; i <= V; i++ {
		if !assigned[i] {
			table = append(table, []int{i})
		}
	}
	for _, elems := range components {
		if elems == nil {
			continue
		}
		sort.Ints(elems)
		table = append(table, elems)
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i][0] < table[j][0]
	})

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	_, _ = fmt.Fprintln(writer, len(table))
	for _, elems := range table {
		for _, e := range elems {
			_, _ = fmt.Fprint(writer, e, " ")
		}
		_, _ = fmt.Fprint(writer, "-1\n")
	}
}
