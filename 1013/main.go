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

func main() {
	T := nextInt()
	for T > 0 {
		signal := nextBytes()
		ok := sub(signal)
		answer := "NO"
		if ok {
			answer = "YES"
		}
		fmt.Println(answer)
		T--
	}
}

var dfa = [9][2]int{
	0: {1, 2},
	1: {8, 0},
	2: {3, 8},
	3: {4, 8},
	4: {4, 5},
	5: {1, 6},
	6: {7, 6},
	7: {4, 0},
	8: {8, 8},
}

func sub(S []byte) bool {
	state := 0
	for _, b := range S {
		dir := int(b - '0')
		state = dfa[state][dir]
	}
	return state == 0 || state == 5 || state == 6
}

// The first try; without knowing about DFA
// func sub(S []byte) bool {
// 	N := len(S)
// 	var left, right func(int) bool
// 	left = func(i int) bool {
// 		if i == N {
// 			return true
// 		}
// 		if !(i < N && S[i] == '1') {
// 			return false
// 		}
// 		i++
// 		if !(i < N && S[i] == '0') {
// 			return false
// 		}
// 		i++
// 		if !(i < N && S[i] == '0') {
// 			return false
// 		}
// 		i++
// 		for i < N && S[i] == '0' {
// 			i++
// 		}
// 		if !(i < N && S[i] == '1') {
// 			return false
// 		}
// 		i++
// 		leftable := false
// 		for i < N && S[i] == '1' {
// 			i++
// 			leftable = true
// 		}
// 		return (leftable && left(i-1)) || right(i)
// 	}
// 	right = func(i int) bool {
// 		if i == N {
// 			return true
// 		}
// 		if !(i < N && S[i] == '0') {
// 			return false
// 		}
// 		i++
// 		if !(i < N && S[i] == '1') {
// 			return false
// 		}
// 		i++
// 		return left(i) || right(i)
// 	}
// 	return left(0) || right(0)
// }
