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
	nextInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		ns := make([]int, len(fields))
		for i, f := range fields {
			ns[i], _ = strconv.Atoi(f)
		}
		return ns
	}

	_ = nextInts()
	h := map[int]bool{}
	for _, x := range nextInts() {
		h[x] = true
	}

	_ = nextInts()
	writer := bufio.NewWriter(os.Stdout)
	for _, x := range nextInts() {
		if h[x] {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}

	_ = writer.Flush()
}
