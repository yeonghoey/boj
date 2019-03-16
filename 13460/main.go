package main

import (
	"bufio"
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

func nextRow() []byte {
	scanner.Scan()
	in := scanner.Bytes()
	bs := make([]byte, len(in))
	copy(bs, in)
	return bs
}

type board [][]byte

func (b board) at(p pos) byte {
	return b[p.r][p.c]
}

type pos struct {
	r, c int
}

func (p pos) add(o pos) pos {
	return pos{p.r + o.r, p.c + o.c}
}

type state struct {
	red, blue pos
}

type node struct {
	state
	count int
}

func main() {
	N := nextInt()
	_ = nextInt()
	B := make(board, N)
	var red, blue pos
	for i := range B {
		row := nextRow()
		for j, b := range row {
			if b == 'R' {
				red = pos{i, j}
			}
			if b == 'B' {
				blue = pos{i, j}
			}
		}
		B[i] = row
	}

	move := func(from, other, dir pos) (to pos, hole bool) {
		to = from
		for {
			if B.at(to) == 'O' {
				hole = true
				break
			}
			next := to.add(dir)
			if B.at(next) != 'O' && next == other {
				break
			}
			if B.at(next) == '#' {
				break
			}
			to = next
		}
		return to, hole
	}

	S := state{red, blue}
	queue := []node{{S, 0}}
	visited := map[state]bool{}
	visited[S] = true
	answer := -1
	for len(queue) > 0 {
		this := queue[0]
		queue = queue[1:]
		count := this.count + 1
		if count > 10 {
			continue
		}
		for _, dir := range []pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			// Try to move red twice to avoid
			// making a bid stuck in the previous position
			red, _ := move(this.red, this.blue, dir)
			blue, bIn := move(this.blue, red, dir)
			red, rIn := move(red, blue, dir)

			if bIn {
				continue
			}

			if rIn {
				answer = count
				break
			}

			next := state{red, blue}
			if visited[next] {
				continue
			}
			visited[next] = true
			queue = append(queue, node{next, count})
		}
		if answer != -1 {
			break
		}
	}
	fmt.Println(answer)
}
