package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N int
	_, _ = fmt.Fscanf(reader, "%d\n", &N)
	I := make([][]byte, N)
	for i := range I {
		s, _ := reader.ReadString('\n')
		I[i] = []byte(strings.TrimSpace(s))
	}
	var compress func(int, int, int) string
	compress = func(y, x int, k int) string {
		first := I[y][x]
		homogeneous := true
		for i := y; i < y+k; i++ {
			for j := x; j < x+k; j++ {
				if I[i][j] != first {
					homogeneous = false
					break
				}
			}
			if !homogeneous {
				break
			}
		}

		if homogeneous {
			return string(first)
		}

		parts := make([]string, 0, 6)
		parts = append(parts, "(")
		parts = append(parts, compress(y, x, k/2))
		parts = append(parts, compress(y, x+k/2, k/2))
		parts = append(parts, compress(y+k/2, x, k/2))
		parts = append(parts, compress(y+k/2, x+k/2, k/2))
		parts = append(parts, ")")
		return strings.Join(parts, "")
	}

	answer := compress(0, 0, N)
	fmt.Println(answer)
}
