package _0211010

//func findTargetSumWays(nums []int, target int) int {
//	res := 0
//	var dfs func(index int, target int)
//	dfs = func(index int, target int) {
//		if index == len(nums) {
//			if target == 0 {
//				res += 1
//			}
//			return
//		}
//		dfs(index+1, target-nums[index])
//		dfs(index+1, target+nums[index])
//	}
//	dfs(0, target)
//	return res
//}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= nums[i] {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	return dp[len(nums)][neg]
}
