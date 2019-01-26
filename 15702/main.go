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

func nextStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	N := nextInt()
	M := nextInt()
	scores := make([]int, N)
	for i := range scores {
		scores[i] = nextInt()
	}

	top, best := -1, -1
	for i := 0; i < M; i++ {
		id := nextInt()
		score := 0
		for _, s := range scores {
			mark := nextStr()
			if mark == "O" {
				score += s
			}
		}
		if (score > best) || (score == best && id < top) {
			top = id
			best = score
		}
	}
	fmt.Printf("%d %d", top, best)
}
