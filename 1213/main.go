package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	odd := 0
	counts := make([]int, 26)
	for _, r := range s {
		idx := int(r - 'A')
		counts[idx]++
		if counts[idx]%2 == 1 {
			odd++
		} else {
			odd--
		}
	}

	answer := "I'm Sorry Hansoo"
	if odd <= 1 {
		n := len(s)
		buf := make([]rune, n)
		pos := 0
		mid := -1
		for i, k := range counts {
			for j := 2; j <= k; j += 2 {
				buf[pos] = rune(i + 'A')
				buf[n-pos-1] = rune(i + 'A')
				pos++
			}
			if k%2 == 1 {
				mid = i
			}
		}
		if mid != -1 {
			buf[n/2] = rune(mid + 'A')
		}
		answer = string(buf)
	}

	fmt.Println(answer)
}
