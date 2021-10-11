package _0211011

import "sort"

func deleteAndEarn(nums []int) int {
	dp := make([]int, len(nums))
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			res = nums[i]
			continue
		}
		ans := 0
		for j := 0; j < i; j++ {
			if nums[j] != nums[i]-1 {
				ans = max(ans, dp[j])
			}
		}
		dp[i] = nums[i] + ans
		res = max(res, dp[i])
	}
	return res
}
