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
	K := nextInt()
	N := nextInt()
	wires := make([]int, K)
	maxLen := 0
	for i := range wires {
		x := nextInt()
		if x > maxLen {
			maxLen = x
		}
		wires[i] = x
	}

	try := func(l int) bool {
		n := 0
		for _, w := range wires {
			n += (w / l)
		}
		return n >= N
	}

	lo, hi := 1, maxLen
	best := 0
	for hi >= lo {
		x := (lo + hi) / 2
		ok := try(x)
		if ok {
			if x > best {
				best = x
			}
			lo = x + 1
		} else {
			hi = x - 1
		}
	}
	fmt.Println(best)
}
