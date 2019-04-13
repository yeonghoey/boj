package main

import (
	"bufio"
	"os"
)

type parser struct {
	digits []byte
	cursor byte
}

func (p *parser) update(d byte) {
	if p.cursor == 4 {
		p.digits = append(p.digits, '0')
	}
	p.digits[len(p.digits)-1] += d * p.cursor
	p.cursor /= 2
	if p.cursor == 0 {
		p.cursor = 4
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)

	p1 := parser{[]byte{'0'}, 1}
	p2 := parser{[]byte{'0'}, 2}
	p4 := parser{[]byte{}, 4}
	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		b := scanner.Bytes()[0]
		if b == '0' || b == '1' {
			d := b - '0'
			p1.update(d)
			p2.update(d)
			p4.update(d)
		} else {
			break
		}

	}

	writer := bufio.NewWriter(os.Stdout)
	var digits []byte
	if p1.cursor == 4 {
		digits = p1.digits
	} else if p2.cursor == 4 {
		digits = p2.digits
	} else if p4.cursor == 4 {
		digits = p4.digits
	}
	for _, b := range digits {
		_ = writer.WriteByte(b)
	}
	_ = writer.Flush()
}
