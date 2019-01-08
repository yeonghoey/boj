package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func nextText() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	N := nextInt()
	M := nextInt()
	nohear := map[string]bool{}
	for i := 0; i < N; i++ {
		s := nextText()
		nohear[s] = true
	}

	nohearsee := []string{}
	for i := 0; i < M; i++ {
		s := nextText()
		if nohear[s] {
			nohearsee = append(nohearsee, s)
		}
	}
	sort.Strings(nohearsee)
	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	_, _ = fmt.Fprintln(writer, len(nohearsee))
	for _, s := range nohearsee {
		_, _ = fmt.Fprintln(writer, s)
	}
}
