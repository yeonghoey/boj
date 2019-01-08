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
	M := nextInt()
	D := make([][]int, N+1)
	for i := range D {
		D[i] = make([]int, M+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			above := D[i-1][j] - D[i-1][j-1]
			left := D[i][j-1]
			this := nextInt()
			D[i][j] = above + left + this
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	K := nextInt()
	for k := 0; k < K; k++ {
		i := nextInt()
		j := nextInt()
		x := nextInt()
		y := nextInt()

		answer := D[x][y] - D[i-1][y] - D[x][j-1] + D[i-1][j-1]
		_, _ = fmt.Fprintln(writer, answer)
	}

}
