package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, L int
	_, _ = fmt.Scanf("%d %d", &N, &L)
	found := false
	for l := L; l <= 100; l++ {
		numer := 2*N - l*(l-1)
		denom := 2 * l
		a1 := numer / denom
		if numer%denom == 0 && a1 >= 0 {
			A := make([]string, l)
			for i := range A {
				A[i] = strconv.Itoa(a1 + i)
			}
			fmt.Println(strings.Join(A, " "))
			found = true
			break
		}
	}

	if !found {
		fmt.Println(-1)
	}
}
