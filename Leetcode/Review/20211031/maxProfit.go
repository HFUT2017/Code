package _0211031

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	if k >= len(prices)/2 {
		return maxProfit1(prices)
	}
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][0] = -prices[0]
		dp[i][1] = 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[j][0] = max(dp[j][0], dp[j-1][1]-prices[i])
			dp[j][1] = max(dp[j][1], dp[j][0]+prices[i])
			res = max(res, max(dp[j][0], dp[j][1]))
		}
	}
	return res
}

func maxProfit1(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][1]
}
