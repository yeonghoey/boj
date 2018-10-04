package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	buf := make([][]rune, n)
	for i := range buf {
		buf[i] = make([]rune, n*2)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n*2; j++ {
			buf[i][j] = ' '
		}
	}

	drawTriangle(n, 0, 0, buf)

	writer := bufio.NewWriter(os.Stdout)
	for i := range buf {
		_, _ = fmt.Fprintln(writer, string(buf[i]))
	}
	_ = writer.Flush()
}

func drawTriangle(n, top, left int, buf [][]rune) {
	if n == 3 {
		copy(buf[top+0][left:left+5], []rune("  *  "))
		copy(buf[top+1][left:left+5], []rune(" * * "))
		copy(buf[top+2][left:left+5], []rune("*****"))
		return
	}

	drawTriangle(n/2, top, left+n/2, buf)
	drawTriangle(n/2, top+n/2, left, buf)
	drawTriangle(n/2, top+n/2, left+n, buf)
}
