package main

import "fmt"

func main() {
	var A, B int
	_, _ = fmt.Scanf("%d %d\n", &A, &B)

	k, nth, sum := 1, 1, 0
	for nth <= B {
		for i := 0; i < k && nth <= B; i++ {
			if nth >= A {
				sum += k
			}
			nth++
		}
		k++
	}
	fmt.Println(sum)
}
