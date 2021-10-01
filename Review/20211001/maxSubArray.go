package _0211001

import "math"

func maxSubArray(nums []int) int {
	res := math.MinInt32
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(dp[i], res)
	}
	return res
}
