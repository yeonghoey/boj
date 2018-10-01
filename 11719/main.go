package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(s))
	}
}
