package main

import "fmt"

func main() {
	var N, F int
	_, _ = fmt.Scanf("%d\n%d", &N, &F)
	n00 := (N / 100) * 100
	r00 := n00 % F
	ans := (F - r00) % F
	fmt.Printf("%02d", ans)
}
