package _0211012

import "sort"

func findNumberOfLIS(nums []int) int {
	dp := [][]int{}
	cnt := [][]int{}
	for _, value := range nums {
		i := sort.Search(len(dp), func(i int) bool {
			return dp[i][len(dp[i])-1] >= value
		})
		c := 1
		if i > 0 {
			k := sort.Search(len(dp[i-1]), func(k int) bool {
				return dp[i-1][k] < value
			})
			c = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
		}

		if i == len(dp) {
			dp = append(dp, []int{value})
			cnt = append(cnt, []int{0, c})
		} else {
			dp[i] = append(dp[i], value)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+c)
		}
	}
	c := cnt[len(cnt)-1]
	return c[len(c)-1]
}
