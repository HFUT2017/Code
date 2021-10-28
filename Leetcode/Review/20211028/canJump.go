package _0211028

func canJump(nums []int) bool {
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > end {
			return false
		}
		end = max(end, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
