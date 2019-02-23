package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

const (
	cleaned = -1
	empty   = 0
	wall    = 1

	north = 0
	east  = 1
	south = 2
	west  = 3
)

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

type pos struct{ r, c int }

var directions = [...]int{north, east, south, west}
var offsets = [...]pos{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (p pos) head(dir int) pos {
	o := offsets[dir]
	return pos{p.r + o.r, p.c + o.c}
}

type robot struct {
	pos
	dir  int
	area [][]int
}

func (r robot) clean() bool {
	if r.area[r.r][r.c] == empty {
		r.area[r.r][r.c] = cleaned
		return true
	}
	return false
}

func (r robot) scan(dir int) int {
	x := r.head(dir)
	return r.area[x.r][x.c]
}

func (r robot) left() int {
	return (r.dir + 3) % 4
}
func (r robot) back() int {
	return (r.dir + 2) % 4
}

func main() {
	N := nextInt()
	M := nextInt()
	r := nextInt()
	c := nextInt()
	d := nextInt()

	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, M)
		for j := range A[i] {
			A[i][j] = nextInt()
		}

	}

	count := 0
	robot := robot{pos{r, c}, d, A}
	for {
		if robot.clean() {
			count++
		}

		moveBack := true
		for _, dir := range directions {
			if robot.scan(dir) == empty {
				moveBack = false
				break
			}
		}
		if moveBack {
			if robot.scan(robot.back()) == wall {
				break
			}
			robot.pos = robot.head(robot.back())
		} else {
			robot.dir = robot.left()
			if robot.scan(robot.dir) == empty {
				robot.pos = robot.head(robot.dir)
			}
		}
	}
	fmt.Println(count)
}
