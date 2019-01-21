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
	var ns []int
	var nMax int
	for {
		n := nextInt()
		if n == 0 {
			break
		}

		if n > nMax {
			nMax = n
		}

		ns = append(ns, n)
	}

	isComposite := makeIsComposite(nMax)

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()
	for _, n := range ns {
		a, b := findAB(isComposite, n)
		if a+b != n {
			_, _ = fmt.Fprintf(writer, "Goldbach's conjecture is wrong.\n")
		} else {
			_, _ = fmt.Fprintf(writer, "%d = %d + %d\n", n, a, b)
		}
	}
}

func makeIsComposite(max int) (isComposite []bool) {
	isComposite = make([]bool, max+1)
	for x := 2; x <= max; x++ {
		if isComposite[x] {
			continue
		}
		for y := x + x; y <= max; y += x {
			isComposite[y] = true
		}
	}
	return
}

func findAB(isComposite []bool, n int) (a, b int) {
	for a = 3; a <= n; a++ {
		b = n - a
		if !isComposite[a] && !isComposite[b] {
			break
		}
	}
	return
}
