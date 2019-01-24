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

func nextRow() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func main() {
	N := nextInt()
	M := nextInt()
	G := make([][]byte, N)
	for i := range G {
		G[i] = nextRow()
	}

	best := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			for side := 1; i+side <= N && j+side <= M; side++ {
				tl := G[i][j]
				tr := G[i][j+side-1]
				bl := G[i+side-1][j]
				br := G[i+side-1][j+side-1]
				if tl == tr && tr == br && br == bl {
					area := side * side
					if area > best {
						best = area
					}
				}
			}
		}
	}

	fmt.Println(best)
}
