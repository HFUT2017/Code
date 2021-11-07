package everyday

func longestSubsequence(arr []int, difference int) int {
	hash := map[int]int{}
	dp := make([]int, len(arr))
	res := 0
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		if hash[arr[i]-difference] != 0 {
			dp[i] = hash[arr[i]-difference] + 1
		}
		hash[arr[i]] = dp[i]
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
