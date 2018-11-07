package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	_, _ = reader.ReadString('\n')
	s, _ := reader.ReadString('\n')
	fs := strings.Fields(s)
	ns := make([]int, len(fs))
	for i, f := range fs {
		ns[i], _ = strconv.Atoi(f)
	}

	current := ns[0]
	best := current
	for _, k := range ns[1:] {
		if current < 0 {
			current = 0
		}
		current += k
		if current > best {
			best = current
		}
	}

	fmt.Println(best)
}
