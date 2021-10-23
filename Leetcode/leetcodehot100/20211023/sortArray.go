package _0211023

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivotIndex := partition(nums, start, end)
	quickSort(nums, start, pivotIndex-1)
	quickSort(nums, pivotIndex+1, end)
}

func partition(nums []int, start, end int) int {
	pivot, left, right := nums[start], start, end
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
	nums[left], nums[start] = nums[start], nums[left]
	return left
}
