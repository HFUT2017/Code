package _0211030

func jump(nums []int) int {
	res := 0
	end := 0
	maxPosition := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}
		if i == end {
			res++
			end = maxPosition
		}
	}
	return res
}
