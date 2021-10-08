package tsearch

func InsertSearch(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (target-nums[left])/(nums[right]-nums[left])*(right-left)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
