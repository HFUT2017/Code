package _0211015

import "sort"

func bestSeqAtIndex(height []int, weight []int) int {
	nums := make([][2]int, len(height))
	for i := 0; i < len(height); i++ {
		nums[i][0] = height[i]
		nums[i][1] = weight[i]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if i == 0 {
			dp[i] = 1
			res = 1
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if nums[i][1] < nums[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
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
