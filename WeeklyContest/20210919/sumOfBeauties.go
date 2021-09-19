package _0210919

func SumOfBeauties(nums []int) int {
	min_value := make([]int, len(nums))
	max_value := make([]int, len(nums))
	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max_value[i] = nums[i]
		} else {
			max_value[i] = max(max_value[i-1], nums[i-1])
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			min_value[i] = nums[i]
		} else {
			min_value[i] = min(min_value[i+1], nums[i+1])
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > max_value[i] && nums[i] < min_value[i] {
			dp[i] = 2
		} else if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			dp[i] = 1
		} else {
			dp[i] = 0
		}
		res += dp[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
