package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func nextInts() []int {
	s, _ := reader.ReadString('\n')
	fs := strings.Fields(s)
	ns := make([]int, len(fs))
	for i, f := range fs {
		ns[i], _ = strconv.Atoi(f)
	}
	return ns
}

func main() {
	_ = nextInts()
	P := nextInts()
	sort.Ints(P)
	tt, sum := 0, 0
	for _, t := range P {
		tt += t
		sum += tt
	}
	fmt.Println(sum)
}
