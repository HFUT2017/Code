package _0210907

// MaxSubArray 给定一个整数数组 nums ，找到一个具有最大和的连续子数组
func MaxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for index, value := range nums {
		if index != 0 {
			dp[index] = max(dp[index-1]+value, value)
		}
		res = max(res, dp[index])
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
