package _0211102

import "math"

func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[i/j]+j)
				dp[i] = min(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}
