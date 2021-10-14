package _0211014

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			res = nums[i]
			continue
		}
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
