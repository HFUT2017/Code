package _0211021

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	tail, cur := 0, 1
	for cur < len(nums) {
		for cur < len(nums) && nums[cur] == nums[tail] {
			cur++
		}
		if cur == len(nums) {
			return tail + 1
		}
		tail++
		nums[tail] = nums[cur]
	}
	return tail + 1
}
