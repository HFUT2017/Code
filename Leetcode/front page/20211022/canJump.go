package _0211022

func canJump(nums []int) bool {
	end := 0
	maxPosition := 0
	for i := 0; i <= maxPosition; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if maxPosition >= len(nums) || i+1 == len(nums) {
			return true
		}
		if i == end {
			end = maxPosition
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
