package tsort

func InsertionSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < len(nums)-1; i++ {
		current := nums[i+1]
		preIndex := i
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex], nums[preIndex+1] = nums[preIndex+1], nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = current
	}
	return
}
