package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)

	x := a
	if (a <= b && b <= c) || (c <= b && b <= a) {
		x = b
	}
	if (a <= c && c <= b) || (b <= c && c <= a) {
		x = c
	}

	fmt.Println(x)
}
