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
func nextBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

type bits int

func (b bits) set(e int) bits {
	b |= (1 << uint(e))
	return b
}

func (b bits) in(e int) bool {
	return (b & (1 << uint(e))) > 0
}

type entries map[bits]map[int]int

func (e entries) update(buyers bits, buyer, price int) {
	last := e[buyers]
	if last == nil {
		last = map[int]int{}
		e[buyers] = last
	}

	current, ok := last[buyer]
	if !ok || (price < current) {
		last[buyer] = price
	}
}

func main() {
	N := nextInt()
	trade := make([][]int, N)
	for i := range trade {
		trade[i] = make([]int, N)
		bs := nextBytes()
		for j := range trade[i] {
			trade[i][j] = int(bs[j] - '0')
		}
	}

	prev := make(entries)
	prev.update(bits(1), 0, 0)
	maxBuyers := 0
	for len(prev) > 0 {
		maxBuyers++
		next := entries{}
		for buyers, last := range prev {
			for i := 0; i < N; i++ {
				if buyers.in(i) {
					continue
				}
				included := buyers.set(i)
				for lastBuyer, lastPrice := range last {
					price := trade[lastBuyer][i]
					if lastPrice <= price {
						next.update(included, i, price)
					}
				}
			}
		}
		prev = next
	}
	fmt.Println(maxBuyers)
}
