package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	names := strings.Split(s, "-")
	ss := make([]byte, len(names))
	for i, name := range names {
		ss[i] = name[0]
	}
	fmt.Println(string(ss))
}
