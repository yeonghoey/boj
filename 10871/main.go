package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, x int
	_, _ = fmt.Fscan(reader, &n, &x)

	var as []string
	for i := 0; i < n; i++ {
		var a int
		_, _ = fmt.Fscan(reader, &a)
		if a < x {
			as = append(as, strconv.Itoa(a))
		}
	}

	fmt.Println(strings.Join(as, " "))
}
