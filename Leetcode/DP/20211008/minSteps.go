package _0211008

import "math"

func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
				dp[i] = min(dp[i], dp[i/j]+j)
			}
		}
	}
	return dp[n]
}
