package _0211009

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		j := i - 1
		ans := 0
		for j >= 0 {
			if nums[j] < nums[i] {
				if dp[j] > ans {
					ans = dp[j]
				}
			}
			j--
		}
		dp[i] = ans + 1
		res = max(res, dp[i])
	}
	return res
}
