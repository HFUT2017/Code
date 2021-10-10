package _0211010

func maxProduct(nums []int) int {
	dp := make([][2]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i][0] = nums[i]
			dp[i][1] = nums[i]
			res = nums[0]
			continue
		}
		dp[i][0] = max(nums[i], max(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
		dp[i][1] = min(nums[i], min(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
		res = max(res, max(dp[i][0], dp[i][1]))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
