package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	output := make([]rune, 0, len(input))
	for _, r := range input {
		if unicode.IsLower(r) {
			output = append(output, unicode.ToUpper(r))
		} else {
			output = append(output, unicode.ToLower(r))
		}
	}
	fmt.Println(string(output))
}
