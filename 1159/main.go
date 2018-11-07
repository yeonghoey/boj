package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var N int
	_, _ = fmt.Fscan(reader, &N)

	counter := map[string]int{}
	for N > 0 {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		counter[s[0:1]]++
		N--
	}

	initials := []string{}
	for it, count := range counter {
		if count >= 5 {
			initials = append(initials, it)
		}
	}

	answer := "PREDAJA"
	if len(initials) > 0 {
		sort.Strings(initials)
		answer = strings.Join(initials, "")
	}
	fmt.Println(answer)
}
