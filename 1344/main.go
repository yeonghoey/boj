package main

import "fmt"

const numTerms = 90 / 5

var nonPrimes = [...]int{0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18}

func main() {
	A := nextP()
	B := nextP()

	C := binomialCoeff()
	PcA := nonPrimeP(A, C)
	PcB := nonPrimeP(B, C)
	fmt.Printf("%.6f\n", 1-PcA*PcB)
}

func nextP() float64 {
	var n int
	_, _ = fmt.Scan(&n)
	return float64(n) / 100.0
}

func binomialCoeff() [][]int {
	C := make([][]int, numTerms+1)
	for i := range C {
		C[i] = make([]int, i+1)
		for j := range C[i] {
			if j == 0 || j == i {
				C[i][j] = 1
				continue
			}
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
	return C
}

func nonPrimeP(p float64, C [][]int) float64 {
	n := numTerms
	psum := 0.0
	for _, k := range nonPrimes {
		pk := float64(C[n][k])
		for i := 0; i < k; i++ {
			pk *= p
		}
		for i := 0; i < n-k; i++ {
			pk *= (1 - p)
		}
		psum += pk

	}
	return psum
}
