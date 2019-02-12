package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
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

func nextByte() byte {
	scanner.Scan()
	return scanner.Bytes()[0]
}

func main() {
	L := nextInt()
	C := nextInt()
	chars := make([]byte, C)
	for i := range chars {
		chars[i] = nextByte()
	}
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	writer := bufio.NewWriter(os.Stdout)
	defer func() { _ = writer.Flush() }()

	buf := make([]byte, L)
	var f func(k, l, consonant, vowel int)
	f = func(k, l, consonant, vowel int) {
		if l == L {
			if consonant >= 2 && vowel >= 1 {
				_, _ = fmt.Fprintln(writer, string(buf))
			}
			return
		}
		for i := k; i < C; i++ {
			buf[l] = chars[i]
			switch chars[i] {
			case 'a', 'e', 'i', 'o', 'u':
				f(i+1, l+1, consonant, vowel+1)
			default:
				f(i+1, l+1, consonant+1, vowel)
			}
		}
	}
	f(0, 0, 0, 0)
}
