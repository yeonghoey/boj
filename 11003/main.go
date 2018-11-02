package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func nextInts() []int {
	s, _ := reader.ReadString('\n')
	fs := strings.Fields(s)
	ns := make([]int, len(fs))
	for i, f := range fs {
		ns[i], _ = strconv.Atoi(f)
	}
	return ns
}

func main() {
	x := nextInts()
	N := x[0]
	L := x[1]
	A := nextInts()
	q := make([]int, 0, N)
	for i := range A {
		r := len(q)
		for r > 0 && q[r-1] > A[i] {
			r--
		}
		q = q[:r]
		q = append(q, A[i])
		if i >= L && q[0] == A[i-L] {
			q = q[1:]
		}
		_, _ = fmt.Fprintf(writer, "%d ", q[0])
	}
	_ = writer.Flush()
}
