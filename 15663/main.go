package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	numbers := make([]int, N)
	for i := range numbers {
		numbers[i] = nextInt()
	}
	sort.Ints(numbers)

	snums := make([]string, N)
	for i, k := range numbers {
		snums[i] = strconv.Itoa(k)
	}

	var buf bytes.Buffer
	selects := make([]int, M)
	output := func() {
		for i, sel := range selects {
			buf.WriteString(snums[sel])
			if i < M-1 {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte('\n')
	}

	used := make([]bool, N)
	var f func(int)
	f = func(m int) {
		if m == M {
			output()
			return
		}
		prev := -1
		for i := 0; i < N; i++ {
			if used[i] || numbers[i] == prev {
				continue
			}
			selects[m] = i
			used[i] = true
			f(m + 1)
			used[i] = false
			prev = numbers[i]
		}
	}

	f(0)
	fmt.Println(buf.String())
}
