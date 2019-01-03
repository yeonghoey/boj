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
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	for {
		k := nextInt()
		if k == 0 {
			break
		}
		a := make([]int, k)
		for i := range a {
			a[i] = nextInt()
		}

		enumerate(a, func(pick []int) {
			line := strings.Trim(fmt.Sprint(pick), "[]")
			_, _ = fmt.Fprintln(writer, line)
		})

		_, _ = fmt.Fprintln(writer)
	}
}

func enumerate(a []int, onPick func([]int)) {
	k := len(a)
	pick := make([]int, 6)

	var f func(int, int)
	f = func(nth, from int) {
		if nth == 6 {
			onPick(pick)
			return
		}
		for i := from; i < k; i++ {
			pick[nth] = a[i]
			f(nth+1, i+1)
		}
	}
	f(0, 0)
}
