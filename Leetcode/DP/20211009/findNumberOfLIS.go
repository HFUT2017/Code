package _0211009

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	res := 0
	maxLen := 0
	for i, x := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			if x > y {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			res = cnt[i]
		} else if dp[i] == maxLen {
			res += cnt[i]
		}
	}
	return res
}
