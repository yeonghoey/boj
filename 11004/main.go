package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	nextInt := func() int {
		_ = scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		return n
	}

	N := nextInt()
	K := nextInt() - 1

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	begin, end := 0, N
	for begin < end {
		i := begin
		pivot := A[begin]
		for j := begin + 1; j < end; j++ {
			if A[j] <= pivot {
				i++
				A[i], A[j] = A[j], A[i]
			}
		}

		A[begin], A[i] = A[i], A[begin]

		if i < K {
			begin = i + 1
			continue
		}

		if i > K {
			end = i
			continue
		}

		break
	}

	fmt.Println(A[K])
}
