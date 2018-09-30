package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	numDwarfs = 9
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	total := 0
	var hs []int
	for i := 0; i < numDwarfs; i++ {
		var h int
		_, _ = fmt.Fscan(reader, &h)
		hs = append(hs, h)
		total += h
	}

	sort.Ints(hs)

	a, b, found := -1, -1, false
	for i := 0; i < numDwarfs-1; i++ {
		for j := i + 1; j < numDwarfs; j++ {
			if hs[i]+hs[j]+100 == total {
				found = true
				a, b = i, j
				break
			}
		}
		if found {
			break
		}
	}

	for i, h := range hs {
		if i == a || i == b {
			continue
		}
		fmt.Println(h)
	}
}
