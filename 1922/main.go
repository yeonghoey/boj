package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type edge struct {
	to, cost int
}

type edgeHeap []*edge

func (eh edgeHeap) Len() int {
	return len(eh)
}

func (eh edgeHeap) Less(i, j int) bool {
	return eh[i].cost < eh[j].cost
}

func (eh edgeHeap) Swap(i, j int) {
	eh[i], eh[j] = eh[j], eh[i]
}

func (eh *edgeHeap) Push(x interface{}) {
	*eh = append(*eh, x.(*edge))
}

func (eh *edgeHeap) Pop() interface{} {
	old := *eh
	n := len(old)
	x := old[n-1]
	*eh = old[0 : n-1]
	return x
}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	N := nextInt()
	M := nextInt()
	G := make(map[int][]edge)
	s := -1
	maxCost := 0
	for i := 0; i < M; i++ {
		a := nextInt()
		b := nextInt()
		c := nextInt()
		G[a] = append(G[a], edge{b, c})
		G[b] = append(G[b], edge{a, c})

		if s == -1 {
			s = a
		}
		if maxCost < c {
			maxCost = c
		}
	}

	eh := &edgeHeap{&edge{s, 0}}
	heap.Init(eh)
	X := map[int]bool{}
	total := 0
	for len(X) < N {
		var pick *edge
		for {
			pick = heap.Pop(eh).(*edge)
			if !X[pick.to] {
				break
			}
		}
		X[pick.to] = true
		total += pick.cost

		for _, e := range G[pick.to] {
			if !X[e.to] {
				ee := e
				heap.Push(eh, &ee)
			}
		}
	}

	fmt.Println(total)
}
