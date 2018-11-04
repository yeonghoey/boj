package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func nextInt() (n int, ok bool) {
	_, err := fmt.Fscan(reader, &n)
	ok = (err == nil)
	return
}

type node struct {
	key   int
	left  *node
	right *node
}

func insert(root *node, key int) {
	var this **node = &root
	for (*this) != nil {
		if key < (*this).key {
			this = &((*this).left)
		} else {
			this = &((*this).right)
		}
	}
	*this = &node{key, nil, nil}
}

func output(this *node) {
	if this == nil {
		return
	}
	output(this.left)
	output(this.right)
	_, _ = fmt.Fprintln(writer, this.key)
}

func main() {
	key, _ := nextInt()
	root := &node{key, nil, nil}
	for {
		key, ok := nextInt()
		if !ok {
			break
		}

		insert(root, key)
	}

	output(root)
	_ = writer.Flush()
}
