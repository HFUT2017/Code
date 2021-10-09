package _0211009

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = 1
			res = 1
			continue
		}
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
