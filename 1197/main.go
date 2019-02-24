package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const maxCost = 2147483647

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

type edge struct{ head, cost int }

type elem struct {
	edge
	index int
}

type elemHeap []*elem

func (h elemHeap) Len() int           { return len(h) }
func (h elemHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h elemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *elemHeap) Push(x interface{}) {
	n := len(*h)
	e := x.(*elem)
	e.index = n
	*h = append(*h, e)
}

func (h *elemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	e := old[n-1]
	e.index = -1
	*h = old[:n-1]
	return e
}

func main() {
	V := nextInt()
	E := nextInt()
	G := make([][]edge, V)
	for i := 0; i < E; i++ {
		// Make 0-based
		a := nextInt() - 1
		b := nextInt() - 1
		cost := nextInt()
		G[a] = append(G[a], edge{b, cost})
		G[b] = append(G[b], edge{a, cost})
	}

	elems := make([]elem, V)
	h := make(elemHeap, 0)
	for i := range elems {
		elems[i] = elem{edge{i, maxCost}, i}
		h = append(h, &(elems[i]))
	}
	elems[0].cost = 0
	heap.Init(&h)
	total := 0
	added := make([]bool, V)
	for len(h) > 0 {
		e := heap.Pop(&h).(*elem)
		added[e.head] = true
		total += e.cost
		for _, ee := range G[e.head] {
			if added[ee.head] {
				continue
			}
			current := &(elems[ee.head])
			if ee.cost < current.cost {
				current.cost = ee.cost
				heap.Fix(&h, current.index)
			}
		}
	}
	fmt.Println(total)
}
