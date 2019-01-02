package main

import "fmt"

func main() {
	var A, B int
	_, _ = fmt.Scanf("%d %d", &A, &B)

	integer := A / B
	remain := A % B
	fractional := make([]byte, 0)
	for i := 0; i < 1000; i++ {
		remain *= 10
		fractional = append(fractional, byte(remain/B+'0'))
		remain = remain % B
		if remain == 0 {
			break
		}
	}
	fmt.Printf("%d.%s\n", integer, string(fractional))
}
