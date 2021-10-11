package _0211011

func LenLongestFibSubseq(arr []int) int {
	dp := make([]int, len(arr))
	hash := map[int]int{}
	res := 0
	for i := 0; i < len(arr); i++ {
		if i < 2 {
			hash[i] = 0
			dp[i] = 0
			continue
		}
		target := arr[i]
		for j := i-1; j >=0; j-- {
			if _, ok := hash[target-arr[j]]; ok {
				if dp[i] == 0 {
					dp[i] = 3
				} else {
					dp[i] = max(dp[i], dp[target-arr[j]]+1)
				}
			}
		}
		hash[arr[i]] = dp[i]
		res = max(res, dp[i])
	}
	for i := 0; i < len(dp); i++ {
		println(dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
