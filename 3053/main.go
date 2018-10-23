package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	_, _ = fmt.Scan(&R)
	fmt.Printf("%.6f\n", R*R*math.Pi)
	fmt.Printf("%.6f\n", R*R*2)
}
