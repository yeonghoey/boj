package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nextInt64 := func() int64 {
		var n int64
		_, err := fmt.Fscan(reader, &n)
		if err != nil {
			panic(err)
		}
		return n
	}

	min := nextInt64()
	max := nextInt64()
	ans := solve(min, max)
	fmt.Println(ans)
}

func solve(min, max int64) int {
	mark := make(map[int64]bool)

	sqrt := int(math.Sqrt(float64(max)))
	sieve := make([]bool, sqrt+1)

	for i := range sieve {
		sieve[i] = true
	}

	for num := 2; num <= sqrt; num++ {
		if !sieve[num] {
			continue
		}

		for i := num + num; i <= sqrt; i += num {
			sieve[i] = false
		}

		square := int64(num) * int64(num)

		from := min - (min % square)
		if from < min {
			from += square
		}

		for i := from; i <= max; i += square {
			if i%square == 0 {
				mark[i] = true
			}
		}
	}

	return int(max-min+1) - len(mark)
}
