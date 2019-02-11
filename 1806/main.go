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

func main() {
	N := nextInt()
	S := nextInt()

	sum := 0
	queue := make([]int, 0, N)
	best := 0
	for i := 0; i < N; i++ {
		x := nextInt()
		sum += x
		queue = append(queue, x)
		for len(queue) > 0 && sum-queue[0] >= S {
			sum -= queue[0]
			queue = queue[1:]
		}
		if sum >= S && (best == 0 || len(queue) < best) {
			best = len(queue)
		}
	}
	fmt.Println(best)
}
