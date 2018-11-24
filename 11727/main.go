package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2]*2 + dp[i-1]
		dp[i] %= 10007
	}
	fmt.Println(dp[n])
}
