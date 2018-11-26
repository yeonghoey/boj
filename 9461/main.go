package main

import "fmt"

func nextInt() int {
	var n int
	_, _ = fmt.Scan(&n)
	return n
}

func main() {
	T := nextInt()
	for T > 0 {
		N := nextInt()
		PN := solve(N)
		fmt.Println(PN)
		T--
	}
}

func solve(N int) int {
	dp := make([]int, N+1)
	base := []int{0, 1, 1, 1, 2, 2}
	copy(dp, base)
	for i := len(base); i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-5]
	}
	return dp[N]
}
