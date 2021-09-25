package tsort

func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		min := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return
}
