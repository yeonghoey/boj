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

type entry struct {
	z, y, x int
	day     int
}

func main() {
	M := nextInt()
	N := nextInt()
	H := nextInt()
	boxes := make([][][]int, H)
	total := 0
	queue := make([]entry, 0)
	for i := range boxes {
		box := make([][]int, N)
		for j := range box {
			row := make([]int, M)
			for k := range row {
				x := nextInt()
				switch x {
				case 0:
					total++
				case 1:
					queue = append(queue, entry{i, j, k, 0})
				}
				row[k] = x
			}
			box[j] = row
		}
		boxes[i] = box
	}

	count := 0
	lastDay := 0
	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		lastDay = e.day
		for _, offset := range []struct{ z, y, x int }{
			{+1, 0, 0}, {-1, 0, 0},
			{0, +1, 0}, {0, -1, 0},
			{0, 0, +1}, {0, 0, -1},
		} {
			z := e.z + offset.z
			y := e.y + offset.y
			x := e.x + offset.x
			if false ||
				!(0 <= z && z < H) ||
				!(0 <= y && y < N) ||
				!(0 <= x && x < M) {
				continue
			}
			if boxes[z][y][x] != 0 {
				continue
			}
			boxes[z][y][x] = 1
			count++
			ee := entry{z, y, x, e.day + 1}
			queue = append(queue, ee)
		}
	}

	answer := -1
	if count == total {
		answer = lastDay
	}
	fmt.Println(answer)
}
