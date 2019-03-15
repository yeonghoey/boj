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

	var answer [10]int
	x, e := N, 1
	for x > 0 {
		q, r := x/10, x%10
		for i := 0; i < 10; i++ {
			answer[i] += q * e
		}
		answer[0] -= e
		for i := 0; i < r; i++ {
			answer[i] += e
		}
		answer[r] += N%e + 1
		x = q
		e *= 10
	}
	line := strings.Trim(fmt.Sprint(answer), "[]")
	fmt.Println(line)
}
