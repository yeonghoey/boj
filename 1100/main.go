package main

import (
	"bufio"
	"fmt"
	"os"
)

const n = 8

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		for j, b := range scanner.Bytes() {
			if (i+j)%2 == 0 && b == 'F' {
				count++
			}
		}
	}
	fmt.Println(count)
}
