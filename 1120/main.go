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

func nextBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func main() {
	A := nextBytes()
	B := nextBytes()
	na := len(A)
	nb := len(B)
	best := nb
	for i := 0; i <= nb-na; i++ {
		count := 0
		for j := 0; j < na; j++ {
			a := A[j]
			b := B[i+j]
			if a != b {
				count++
			}
		}
		if count < best {
			best = count
		}
	}
	fmt.Println(best)
}
