package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

type time int

func nextTime() time {
	scanner.Scan()
	bs := scanner.Bytes()
	h := int(bs[0]-'0')*10 + int(bs[1]-'0')
	m := int(bs[3]-'0')*10 + int(bs[4]-'0')
	s := int(bs[6]-'0')*10 + int(bs[7]-'0')
	return time(h*3600 + m*60 + s)
}

func (t time) format() string {
	x := t
	h, x := x/3600, x%3600
	m, s := x/60, x%60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func main() {
	now := nextTime()
	start := nextTime()
	if now > start {
		start += 24 * 3600
	}
	fmt.Println((start - now).format())
}
