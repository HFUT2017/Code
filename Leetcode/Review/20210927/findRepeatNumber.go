package _0210927

func findRepeatNumber(nums []int) int {
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return nums[i]
		}
		hash[nums[i]] = true
	}
	return 0
}
