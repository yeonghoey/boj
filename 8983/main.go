package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	M := nextInt()
	N := nextInt()
	L := nextInt()

	slots := make([]int, M)
	for i := range slots {
		slots[i] = nextInt()
	}
	sort.Ints(slots)

	type animal struct{ x, y int }
	animals := make([]animal, N)
	for i := range animals {
		x := nextInt()
		y := nextInt()
		animals[i] = animal{x, y}
	}
	sort.Slice(animals, func(i, j int) bool {
		if animals[i].x == animals[j].x {
			return animals[i].y < animals[j].y
		}
		return animals[i].x < animals[j].x
	})

	si := 0
	count := 0
	for _, a := range animals {
		for si < M-1 && abs(slots[si]-a.x) > abs(slots[si+1]-a.x) {
			si++
		}
		dist := abs(slots[si]-a.x) + a.y
		if dist <= L {
			count++
		}
	}
	fmt.Println(count)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
