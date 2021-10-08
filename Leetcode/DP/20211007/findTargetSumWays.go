package _0211007

////回溯
//func findTargetSumWays(nums []int, target int) int {
//	res:=0
//	var dfs func(sum int,index int)
//	dfs=func(sum int,index int){
//		if index>=len(nums){
//			if sum==target{
//				res+=1
//			}
//			return
//		}
//		dfs(sum+nums[index],index+1)
//		dfs(sum-nums[index],index+1)
//	}
//	dfs(0,0)
//	return res
//}

//动态规划
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, value := range nums {
		sum += value
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
