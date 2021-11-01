package _0211101

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i*i <= n; i++ {
		dp[i*i] = 1
	}
	for i := 1; i <= n; i++ {
		if dp[i] != 0 {
			continue
		}
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
