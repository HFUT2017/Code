package _0211010

func canPartition(nums []int) bool {
	dp := make([][]bool, len(nums)+1)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum/2+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	for j := 1; j <= sum/2; j++ {
		for i := 1; i <= len(nums); i++ {
			if nums[i-1] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum/2]
}