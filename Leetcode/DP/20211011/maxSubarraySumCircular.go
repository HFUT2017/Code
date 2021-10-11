package _0211011

import "math"

func maxSubarraySumCircular(nums []int) int {
	dp := make([]int, len(nums)+1)
	dpm := make([]int, len(nums)+1)
	min_, max_, sum := math.MaxInt32, math.MinInt32, 0
	for i := 1; i <= len(nums); i++ {
		sum += nums[i-1]
		dp[i] = max(nums[i-1], dp[i-1]+nums[i-1])
		dpm[i] = min(nums[i-1], dpm[i-1]+nums[i-1])
		max_ = max(max_, dp[i])
		min_ = min(min_, dpm[i])
	}
	if sum-min_ == 0 || max_ > sum-min_ {
		return max_
	} else {
		return sum - min_
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
