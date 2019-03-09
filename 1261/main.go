package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
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

func nextBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

type pos struct {
	y, x int
}

type trace struct {
	pos
	count int
	index int
}

type traceHeap []*trace

func (th traceHeap) Len() int {
	return len(th)
}

func (th traceHeap) Less(i, j int) bool {
	return th[i].count < th[j].count
}

func (th traceHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
	th[i].index = i
	th[j].index = j
}

func (th *traceHeap) Push(x interface{}) {
	n := len(*th)
	t := x.(*trace)
	t.index = n
	*th = append(*th, t)
}

func (th *traceHeap) Pop() interface{} {
	old := *th
	n := len(old)
	t := old[n-1]
	t.index = -1
	*th = old[0 : n-1]
	return t
}

func main() {
	M := nextInt()
	N := nextInt()
	maze := make([][]byte, N)
	for i := range maze {
		maze[i] = make([]byte, M)
		copy(maze[i], nextBytes())
	}

	start := &trace{pos{0, 0}, 0, -1}

	traces := make([][]*trace, N)
	for i := range traces {
		traces[i] = make([]*trace, M)
	}
	traces[0][0] = start
	traceHeap := make(traceHeap, 0, N*M)
	heap.Push(&traceHeap, start)

	var answer int
	for len(traceHeap) > 0 {
		this := heap.Pop(&traceHeap).(*trace)
		if this.y == N-1 && this.x == M-1 {
			answer = this.count
			break
		}
		for _, d := range []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			p := pos{this.y + d.y, this.x + d.x}
			if !(0 <= p.y && p.y < N) || !(0 <= p.x && p.x < M) {
				continue
			}
			count := this.count
			if maze[p.y][p.x] == '1' {
				count++
			}

			next := traces[p.y][p.x]
			if next != nil {
				if count < next.count {
					next.count = count
					heap.Fix(&traceHeap, next.index)
				}
			} else {
				next = &trace{p, count, -1}
				traces[p.y][p.x] = next
				heap.Push(&traceHeap, next)
			}
		}
	}
	fmt.Println(answer)
}
