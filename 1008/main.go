package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	A, B := nextFloat(), nextFloat()
	fmt.Printf("%.9f", A/B)
}

func nextFloat() float64 {
	var f float64
	_, err := fmt.Fscan(reader, &f)
	if err != nil {
		panic(err)
	}
	return f
}
