package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		ns := make([]int, len(fields))
		for i, f := range fields {
			ns[i], _ = strconv.Atoi(f)
		}
		return ns
	}

	dimA := nextInts()
	N, M := dimA[0], dimA[1]
	A := make([][]int, N)
	for i := range A {
		A[i] = nextInts()
	}

	dimB := nextInts()
	M, K := dimB[0], dimB[1]
	B := make([][]int, M)
	for i := range B {
		B[i] = nextInts()
	}

	C := make([][]int, N)
	for i := range C {
		C[i] = make([]int, K)
	}

	for ni := 0; ni < N; ni++ {
		for ki := 0; ki < K; ki++ {
			for mi := 0; mi < M; mi++ {
				C[ni][ki] += A[ni][mi] * B[mi][ki]
			}
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	fields := make([]string, K)
	for _, row := range C {
		for i, x := range row {
			fields[i] = strconv.Itoa(x)
		}
		fmt.Fprintln(writer, strings.Join(fields, " "))
	}
	_ = writer.Flush()
}
