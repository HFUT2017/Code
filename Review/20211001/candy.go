package _0211001

func candy(ratings []int) int {
	dp := make([]int, len(ratings))
	res := 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	right := 0
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right = right + 1
		} else {
			right = 1
		}
		res += max(dp[i], right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
