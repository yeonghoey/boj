package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(nil, 100002)
}

func nextInt() int {
	scanner.Scan()
	sign, x := 1, 0
	for _, b := range scanner.Bytes() {
		if b == '-' {
			sign = -1
			continue
		}
		x *= 10
		x += int(b - '0')
	}
	return sign * x
}

func nextBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

type char struct {
	prev, next *char
	b          byte
}

func (c *char) left() *char {
	if c.prev == nil {
		return c
	}
	return c.prev
}

func (c *char) right() *char {
	if c.next == nil {
		return c
	}
	return c.next
}

func (c *char) remove() *char {
	if c.prev == nil {
		return c
	}
	cp := c.prev
	cpp := cp.prev
	if cpp != nil {
		cpp.next = c
	}
	c.prev = cpp
	return c
}

func (c *char) insert(b byte) *char {
	cp := c.prev
	cnew := &char{cp, c, b}
	if cp != nil {
		cp.next = cnew
	}
	c.prev = cnew
	return c
}

func (c *char) head() *char {
	x := c
	for x.prev != nil {
		x = x.prev
	}
	return x
}

func main() {
	init := nextBytes()
	init = append(init, '\n')
	prev := &char{nil, nil, init[0]}
	cursor := prev
	for _, b := range init[1:] {
		cursor = &char{prev, nil, b}
		prev.next = cursor
		prev = cursor
	}

	N := nextInt()
	for i := 0; i < N; i++ {
		line := nextBytes()
		op := line[0]
		if op == 'L' {
			cursor = cursor.left()
		}

		if op == 'D' {
			cursor = cursor.right()
		}

		if op == 'B' {
			cursor = cursor.remove()
		}

		if op == 'P' {
			x := line[2]
			cursor = cursor.insert(x)
		}
	}

	cursor = cursor.head()
	line := []byte{}
	for cursor.next != nil {
		line = append(line, cursor.b)
		cursor = cursor.next
	}

	fmt.Println(string(line))
}
