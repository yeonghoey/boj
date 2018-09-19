package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	x int
	y int
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	t := nextInt()
	for t > 0 {
		nextInt() // m
		nextInt() // n
		k := nextInt()
		poss := make(map[pos]bool)
		for i := 0; i < k; i++ {
			x := nextInt()
			y := nextInt()
			poss[pos{x, y}] = true
		}
		count := solve(poss)
		fmt.Println(count)
		t--
	}
}

func solve(poss map[pos]bool) int {
	count, mark := 0, make(map[pos]bool)
	for p := range poss {
		if !mark[p] && poss[p] {
			bfs(p, mark, poss)
			count++
		}
	}
	return count
}

func bfs(p pos, mark, poss map[pos]bool) {
	q := []pos{p}

	for h := 0; h < len(q); h++ {
		cur := q[h]

		if !poss[cur] || mark[cur] {
			continue
		}

		mark[cur] = true

		q = append(q,
			pos{cur.x - 1, cur.y},
			pos{cur.x + 1, cur.y},
			pos{cur.x, cur.y - 1},
			pos{cur.x, cur.y + 1})
	}
}

func nextInt() int {
	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		panic(err)
	}
	return n
}
