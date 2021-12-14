package Designandanalysisofalgorithms

func MinMax(nums []int) (int, int) {
	if len(nums) == 1 {
		return nums[0], nums[0]
	}
	leftMin, leftMax := MinMax(nums[:len(nums)/2])
	rightMin, rightMax := MinMax(nums[len(nums)/2:])
	return min(leftMin, rightMin), max(leftMax, rightMax)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
