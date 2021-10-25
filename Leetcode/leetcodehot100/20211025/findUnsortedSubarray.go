package _0211025

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	leftMax, rightMin := nums[0], nums[len(nums)-1]
	left, right := 0, -1
	for i := 0; i < len(nums); i++ {
		if leftMax > nums[i] {
			right = i
		} else {
			leftMax = nums[i]
		}
		if rightMin < nums[len(nums)-i-1] {
			left = len(nums) - i - 1
		} else {
			rightMin = nums[len(nums)-i-1]
		}
	}
	return right - left + 1
}
