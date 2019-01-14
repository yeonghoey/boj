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

func main() {
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(len(s))
}
