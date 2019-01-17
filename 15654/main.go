package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

func main() {
	N := nextInt()
	M := nextInt()
	numbers := make([]int, N)
	used := make([]bool, N)
	selects := make([]int, M)

	for i := range numbers {
		numbers[i] = nextInt()
	}
	sort.Ints(numbers)

	var buf bytes.Buffer
	var f func(int, int)
	f = func(n, m int) {
		if m == M {
			s := strings.Trim(fmt.Sprint(selects), "[]")
			fmt.Fprintln(&buf, s)
			return
		}
		for i := 0; i < N; i++ {
			if used[i] {
				continue
			}

			selects[m] = numbers[i]
			used[i] = true
			f(i+1, m+1)
			used[i] = false
		}
	}

	f(0, 0)
	fmt.Println(buf.String())
}
