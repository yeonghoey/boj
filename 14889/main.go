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
	S := make([][]int, N)
	for i := range S {
		S[i] = make([]int, N)
		for j := range S[i] {
			S[i][j] = nextInt()
		}
	}

	calcScore := func(team []int) int {
		score := 0
		for i := 0; i < len(team); i++ {
			ti := team[i]
			for j := i + 1; j < len(team); j++ {
				tj := team[j]
				score += S[ti][tj]
				score += S[tj][ti]
			}
		}
		return score
	}

	pickOther := func(team []int) []int {
		ti, oi := 0, 0
		other := make([]int, N/2)
		for i := 0; i < N; i++ {
			if ti < N/2 && team[ti] == i {
				ti++
				continue
			}
			other[oi] = i
			oi++
		}
		return other
	}

	team := make([]int, N/2)
	best := 10000 // Possible maximum: 10 x 9 x 100
	var f func(int, int)
	f = func(num, k int) {
		if num == N/2 {
			teamScore := calcScore(team)
			other := pickOther(team)
			otherScore := calcScore(other)
			diff := teamScore - otherScore
			if diff < 0 {
				diff = -diff
			}
			if diff < best {
				best = diff
			}
			return
		}
		for i := k; i < N; i++ {
			team[num] = i
			f(num+1, i+1)
		}
	}
	f(0, 0)
	fmt.Println(best)
}
