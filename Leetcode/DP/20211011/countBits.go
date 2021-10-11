package _0211011

func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i *= 2 {
		dp[i] = 1
	}
	for i := 1; i <= n; i++ {
		if dp[i] == 0 {
			if i%2 == 0 {
				dp[i] = dp[i/2]
			} else {
				dp[i] = dp[i-1] + 1
			}
		}
	}
	return dp
}
