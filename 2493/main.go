package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type tower struct {
	number int
	height int
}

func main() {
	N := nextInt()
	S := []tower{{0, 0}}
	A := []int{}
	for i := 1; i <= N; i++ {
		height := nextInt()
		for len(S) > 1 && S[len(S)-1].height < height {
			S = S[:len(S)-1]
		}
		A = append(A, S[len(S)-1].number)
		S = append(S, tower{i, height})
	}
	s := strings.Trim(fmt.Sprint(A), "[]")
	fmt.Println(s)
}
