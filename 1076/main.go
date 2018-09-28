package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nextString := func() string {
		var s string
		_, err := fmt.Fscan(reader, &s)
		if err != nil {
			panic(err)
		}
		return s
	}

	table := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	a := table[nextString()]
	b := table[nextString()]
	m := table[nextString()]

	x := int64(a*10 + b)
	for i := 0; i < m; i++ {
		x *= 10
	}

	fmt.Println(x)
}
