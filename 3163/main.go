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
	T := nextInt()
	for T > 0 {
		sub()
		T--
	}
}

type ant struct{ a, time int }

func sub() {
	N := nextInt()
	L := nextInt()
	k := nextInt()
	ants := make([]ant, N)
	lants := make([]ant, 0, N)
	rants := make([]ant, 0, N)
	for i := range ants {
		p := nextInt()
		a := nextInt()
		x := ant{a, L - p}
		if a < 0 {
			x.time = p
			lants = append(lants, x)
		} else {
			rants = append(rants, x)
		}
		ants[i] = ant{a, 0}
	}
	// ants provided as p_i < p_{i+1}
	for i, x := range lants {
		ants[i].time = x.time
	}
	for i, x := range rants {
		ants[len(lants)+i].time = x.time
	}
	sort.Slice(ants, func(i, j int) bool {
		it := ants[i].time
		jt := ants[j].time
		if it == jt {
			return ants[i].a < ants[j].a
		}
		return it < jt
	})
	fmt.Println(ants[k-1].a)
}
