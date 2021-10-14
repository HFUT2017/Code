package _0211014

func waysToChange(n int) int {
	dp := make([]int, n+1)
	coins := []int{1, 5, 10, 25}
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= n; j++ {
			if j >= coins[i] {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % 1000000007
}
