package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct{ y, x int }

type trace struct {
	pos
	day int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		ns := make([]int, len(fields))
		for i, f := range fields {
			ns[i], _ = strconv.Atoi(f)
		}
		return ns
	}

	MN := nextInts()
	M, N := MN[0], MN[1]

	box := make([][]int, N)
	for i := range box {
		box[i] = nextInts()
	}

	queue := make([]trace, 0, N*M)
	for i, row := range box {
		for j := range row {
			x := row[j]
			if x == 1 {
				queue = append(queue, trace{pos{i, j}, 0})
			}
		}
	}

	offsets := [...]pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	maxDay := 0
	for len(queue) > 0 {
		this := queue[0]
		queue = queue[1:]

		if this.day > maxDay {
			maxDay = this.day
		}

		for _, os := range offsets {
			next := pos{this.y + os.y, this.x + os.x}
			if next.y < 0 || next.y >= N || next.x < 0 || next.x >= M {
				continue
			}

			if box[next.y][next.x] != 0 {
				continue
			}

			box[next.y][next.x] = 1
			queue = append(queue, trace{next, this.day + 1})
		}
	}

	answer := maxDay

	for _, row := range box {
		for j := range row {
			if row[j] == 0 {
				answer = -1
				break
			}
		}
		if answer == -1 {
			break
		}
	}

	fmt.Println(answer)
}
