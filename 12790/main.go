package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner
var writer *bufio.Writer

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
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

func writeInt(x int) {
	_, _ = fmt.Fprintln(writer, x)
}

func main() {
	T := nextInt()
	for T > 0 {
		sub()
		T--
	}
	_ = writer.Flush()
}

func sub() {
	// Base
	HP := nextInt()
	MP := nextInt()
	ATT := nextInt()
	DEF := nextInt()

	// Equipment
	HP += nextInt()
	MP += nextInt()
	ATT += nextInt()
	DEF += nextInt()

	power := 1*max(HP, 1) + 5*max(MP, 1) + 2*max(ATT, 0) + 2*DEF
	writeInt(power)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
