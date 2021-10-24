package _0211024

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		for right < len(nums) && nums[right] == 0 {
			right++
		}
		if right == len(nums) {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right++
	}
}
