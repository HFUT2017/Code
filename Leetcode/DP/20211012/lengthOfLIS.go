package _0211012

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, value := range nums {
		if len(dp) == 0 || value > dp[len(dp)-1] {
			dp = append(dp, value)
		} else {
			l, r := 0, len(dp)-1
			pos := 0
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= value {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = value
		}
	}
	return len(dp)
}
