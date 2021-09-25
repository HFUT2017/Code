package tsort

func QuickSort(nums []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(nums, startIndex, endIndex)
	QuickSort(nums, startIndex, pivotIndex-1)
	QuickSort(nums, pivotIndex+1, endIndex)
}

func partition(nums []int, startIndex, endIndex int) int {
	pivot, left, right := nums[startIndex], startIndex, endIndex
	for left != right {
		for left < right && nums[right] > pivot {
			right--
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[startIndex], nums[left] = nums[left], nums[startIndex]
	return left
}
