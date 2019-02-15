package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Bytes()
	count := make([]int, 26)
	for _, b := range word {
		count[int(b-'a')]++
	}
	s := fmt.Sprint(count)
	line := strings.Trim(s, "[]")
	fmt.Println(line)
}
