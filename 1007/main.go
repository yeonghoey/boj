package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	y, x int64
}

func (a point) add(b point) point {
	return point{a.y + b.y, a.x + b.x}
}

func (a point) sub(b point) point {
	return point{a.y - b.y, a.x - b.x}
}

func (a point) mag2() int64 {
	return a.y*a.y + a.x*a.x
}

func main() {
	nextInt := newNextInt()
	T := nextInt()
	for T > 0 {
		process(nextInt)
		T--
	}
}

func process(nextInt func() int) {
	N := nextInt()
	points := make([]point, N)

	for i := range points {
		y := nextInt()
		x := nextInt()
		points[i] = point{int64(y), int64(x)}
	}

	picked := make([]bool, N)
	var answer2 int64 = (2000000 * 2000000) * 20
	var walk func(int, int)

	walk = func(x, remain int) {
		if remain == 0 {
			total := point{0, 0}
			for i, p := range points {
				if picked[i] {
					total = total.add(p)
				} else {
					total = total.sub(p)
				}
			}
			m2 := total.mag2()
			if m2 < answer2 {
				answer2 = m2
			}
			return
		}
		for i := x; i < N; i++ {
			picked[i] = true
			walk(i+1, remain-1)
			picked[i] = false
		}
	}

	walk(0, N/2)
	answer := math.Sqrt(float64(answer2))
	fmt.Printf("%.6f\n", answer)
}

func newNextInt() func() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return func() int {
		scanner.Scan()
		x, neg := 0, false
		for _, b := range scanner.Bytes() {
			if b == '-' {
				neg = true
				continue
			}
			x *= 10
			x += int(b - '0')
		}
		if neg {
			return -x
		}
		return x
	}
}
