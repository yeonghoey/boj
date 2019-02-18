package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	maxCost = 1000 * 100000
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

type bus struct {
	to, cost int
	index    int
}

type busHeap []*bus

func (h busHeap) Len() int {
	return len(h)
}

func (h busHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h busHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *busHeap) Push(x interface{}) {
	n := len(*h)
	b := x.(*bus)
	b.index = n
	*h = append(*h, b)
}

func (h *busHeap) Pop() interface{} {
	old := *h
	n := len(old)
	b := old[n-1]
	b.index = -1
	*h = old[:n-1]
	return b
}

func (h *busHeap) update(b *bus, cost int) {
	if cost < b.cost {
		b.cost = cost
		heap.Fix(h, b.index)
	}
}

func main() {
	N := nextInt()
	M := nextInt()
	G := make([][]bus, N+1)
	for i := 0; i < M; i++ {
		from := nextInt()
		to := nextInt()
		cost := nextInt()
		G[from] = append(G[from], bus{to, cost, -1})
	}

	S := nextInt()
	E := nextInt()

	buses := make([]bus, N+1)
	h := make(busHeap, 0, N)
	for i := 1; i <= N; i++ {
		buses[i].to = i
		buses[i].cost = maxCost
		buses[i].index = len(h)
		h = append(h, &buses[i])
	}
	buses[S].cost = 0
	heap.Init(&h)
	visited := make([]bool, N+1)
	answer := -1
	for len(h) > 0 {
		x := heap.Pop(&h).(*bus)
		if x.to == E {
			answer = x.cost
			break
		}
		visited[x.to] = true
		for _, b := range G[x.to] {
			if !visited[b.to] {
				cost := x.cost + b.cost
				h.update(&buses[b.to], cost)
			}
		}
	}
	fmt.Println(answer)
}
