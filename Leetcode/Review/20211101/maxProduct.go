package _0211101

func maxProduct(nums []int) int {
	max_dp := make([]int, len(nums))
	min_dp := make([]int, len(nums))
	max_dp[0] = nums[0]
	min_dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		max_dp[i] = max(max_dp[i-1]*nums[i], max(min_dp[i-1]*nums[i], nums[i]))
		min_dp[i] = min(min_dp[i-1]*nums[i], min(max_dp[i-1]*nums[i], nums[i]))
		res = max(res, max_dp[i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
