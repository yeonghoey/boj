package main

import (
	"bufio"
	"fmt"
	"os"
)

type elem struct {
	x      int
	rank   int
	parent *elem
}

type uniFind map[int]*elem

func (uf uniFind) get(x int) *elem {
	e, exists := uf[x]
	if !exists {
		e = &elem{x, 0, nil}
		uf[x] = e
	}
	return e
}

func (uf uniFind) union(a, b int) {
	leaderA := uf.find(a)
	leaderB := uf.find(b)

	// There is no need to union
	if leaderA == leaderB {
		return
	}

	if leaderA.rank == leaderB.rank {
		leaderA.rank++
	}
	if leaderA.rank < leaderB.rank {
		leaderA.parent = leaderB
	} else {
		leaderB.parent = leaderA
	}
}

func (uf uniFind) find(x int) *elem {
	this := uf.get(x)

	// Find the leader.
	trace := []*elem{}
	for this.parent != nil {
		trace = append(trace, this)
		this = this.parent
	}

	// Update the parents of the elements on the path to the leader.
	for _, e := range trace {
		e.parent = this
	}

	return this
}

var scanner *bufio.Scanner
var writer *bufio.Writer

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
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
	defer func() { _ = writer.Flush() }()

	_ = nextInt()
	m := nextInt()

	uf := uniFind{}
	for i := 0; i < m; i++ {
		op := nextInt()
		a := nextInt()
		b := nextInt()

		if op == 0 {
			uf.union(a, b)
		}

		if op == 1 {
			response := "NO"
			if uf.find(a) == uf.find(b) {
				response = "YES"
			}
			_, _ = fmt.Fprintln(writer, response)
		}
	}
}
