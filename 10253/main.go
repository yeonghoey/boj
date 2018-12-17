package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	T := nextInt()
	for T > 0 {
		a := nextInt()
		b := nextInt()
		x := solve(a, b)
		fmt.Println(x)
		T--
	}

}

func solve(a, b int) int {
	var x int
	for a != 0 {
		d := math.Ceil(float64(b) / float64(a))
		x = int(d)
		a = a*x - b
		b = b * x
	}
	return x
}
