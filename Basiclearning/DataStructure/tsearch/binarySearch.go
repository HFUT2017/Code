package tsearch

import (
	"sort"
)

func BinarySearch(nums []int, target int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearch_(nums []int, target int, left, right int) int {
	if left > right {
		return -1
	}
	sort.Ints(nums)
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return BinarySearch_(nums, target, mid+1, right)
	} else {
		return BinarySearch_(nums, target, left, mid-1)
	}
}
