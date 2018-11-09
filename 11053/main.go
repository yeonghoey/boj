package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	N := nextInts()[0]
	A := nextInts()
	D := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		D[i] = 1
		for j := i + 1; j < N; j++ {
			if A[i] < A[j] && D[i] < D[j]+1 {
				D[i] = D[j] + 1
			}
		}
	}

	best := D[0]
	for _, l := range D[1:] {
		if l > best {
			best = l
		}
	}

	fmt.Println(best)
}

func nextInts() []int {
	s, _ := reader.ReadString('\n')
	fields := strings.Fields(s)
	ints := make([]int, len(fields))
	for i, field := range fields {
		ints[i], _ = strconv.Atoi(field)
	}
	return ints
}
