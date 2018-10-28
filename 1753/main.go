package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex = int

type graph map[vertex]map[vertex]int

type entry struct {
	tail  vertex
	score int
	index int
}

type priorityQueue struct {
	entries []*entry
	lookup  map[vertex]*entry
}

func newPQ() *priorityQueue {
	return &priorityQueue{[]*entry{}, map[vertex]*entry{}}
}

func (pq priorityQueue) Len() int { return len(pq.entries) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq.entries[i].score < pq.entries[j].score
}

func (pq priorityQueue) Swap(i, j int) {
	pq.entries[i], pq.entries[j] = pq.entries[j], pq.entries[i]
	pq.entries[i].index = i
	pq.entries[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(pq.entries)
	e := x.(*entry)
	e.index = n
	pq.entries = append(pq.entries, e)
	pq.lookup[e.tail] = e
}

func (pq *priorityQueue) Pop() interface{} {
	old := pq.entries
	n := len(old)
	e := old[n-1]
	e.index = -1
	pq.entries = old[0 : n-1]
	delete(pq.lookup, e.tail)
	return e
}

func (pq *priorityQueue) upsert(e *entry) {
	ee, ok := pq.lookup[e.tail]
	if ok {
		// Update
		if e.score < ee.score {
			ee.score = e.score
			heap.Fix(pq, ee.index)
		}
	} else {
		// Insert
		heap.Push(pq, e)
	}
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

	VE := nextInts()
	V, E := VE[0], VE[1]
	K := nextInts()[0]
	G := make(graph)
	for i := 0; i < E; i++ {
		e := nextInts()
		u, v, w := e[0], e[1], e[2]
		if G[u] == nil {
			G[u] = make(map[vertex]int)
		}

		w1, ok := G[u][v]
		if ok && w1 < w {
			continue
		}

		G[u][v] = w
	}

	shortests := dijkstra(V, E, K, G)

	writer := bufio.NewWriter(os.Stdout)
	for v := 1; v <= V; v++ {
		dist, ok := shortests[v]
		if ok {
			_, _ = fmt.Fprintln(writer, dist)
		} else {
			_, _ = fmt.Fprintln(writer, "INF")
		}
	}
	_ = writer.Flush()
}

func dijkstra(V, E int, K vertex, G graph) map[vertex]int {
	shortests := map[vertex]int{}
	pq := newPQ()
	heap.Push(pq, &entry{K, 0, -1})
	for pq.Len() > 0 {
		e := heap.Pop(pq).(*entry)
		shortests[e.tail] = e.score
		for head, weight := range G[e.tail] {
			if _, ok := shortests[head]; !ok {
				pq.upsert(&entry{head, e.score + weight, -1})
			}
		}
	}
	return shortests
}
