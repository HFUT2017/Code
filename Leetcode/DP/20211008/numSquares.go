package _0211008

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i*i <= n; i++ {
		dp[i*i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j*j < i; j++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
