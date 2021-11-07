package everyday

func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == i {
			continue
		}
		for nums[i] < n && nums[nums[i]] != nums[i] {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}
