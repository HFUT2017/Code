package _0211014

func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	index3, index5, index7 := 0, 0, 0
	for i := 1; i < k; i++ {
		dp[i] = min(dp[index3]*3, min(dp[index5]*5, dp[index7]*7))
		if dp[i] == dp[index3]*3 {
			index3++
		}
		if dp[i] == dp[index5]*5 {
			index5++
		}
		if dp[i] == dp[index7]*7 {
			index7++
		}
	}
	return dp[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
