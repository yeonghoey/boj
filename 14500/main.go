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

type pos struct{ r, c int }

var tetrominos = [19][4]pos{
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
	{{0, 0}, {-1, 0}, {-1, 1}, {-2, 1}},
	{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
	{{0, 0}, {0, 1}, {-1, 1}, {-1, 2}},
	{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
	{{0, 0}, {-1, 0}, {-2, 0}, {-2, 1}},
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}},
	{{0, 0}, {0, 1}, {-1, 1}, {-2, 1}},
	{{0, 0}, {-1, 0}, {-1, 1}, {-1, 2}},
	{{0, 0}, {0, 1}, {0, 2}, {1, 2}},
	{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
	{{0, 0}, {0, 1}, {0, 2}, {-1, 2}},
	{{0, 0}, {-1, 0}, {1, 0}, {0, 1}},
	{{0, 0}, {-1, 1}, {0, 1}, {1, 1}},
	{{0, 0}, {0, 1}, {0, 2}, {1, 1}},
	{{0, 0}, {0, 1}, {0, 2}, {-1, 1}},
}

func main() {
	N := nextInt()
	M := nextInt()
	paper := make([][]int, N)
	for i := range paper {
		paper[i] = make([]int, M)
		for j := range paper[i] {
			paper[i][j] = nextInt()
		}
	}

	try := func(r, c int) int {
		best := 0
		for _, t := range tetrominos {
			sum := 0
			for _, p := range t {
				rr := r + p.r
				cc := c + p.c
				if !(0 <= rr && rr < N) || !(0 <= cc && cc < M) {
					sum = -1
					break
				}
				sum += paper[rr][cc]
			}
			if sum > best {
				best = sum
			}
		}
		return best
	}

	answer := 0
	for i := range paper {
		for j := range paper[i] {
			x := try(i, j)
			if x > answer {
				answer = x
			}
		}
	}
	fmt.Println(answer)
}
