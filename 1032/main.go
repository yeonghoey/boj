package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		panic(err)
	}

	var ss []string

	for i := 0; i < n; i++ {
		var s string
		_, err := fmt.Fscan(reader, &s)
		if err != nil {
			panic(err)
		}

		ss = append(ss, s)
	}

	var ans []byte

	m := len(ss[0])
	for i := 0; i < m; i++ {
		r := ss[0][i]
		useQ := false

		for j := 1; j < n; j++ {
			if r != ss[j][i] {
				useQ = true
				break
			}
		}

		if useQ {
			ans = append(ans, '?')
		} else {
			ans = append(ans, r)
		}
	}

	fmt.Println(string(ans))
}
