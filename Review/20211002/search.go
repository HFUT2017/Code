package _0211002

func search(nums []int, target int) int {
	var binarySearch func(nums []int, left, right int, target int) int
	binarySearch = func(nums []int, left, right int, target int) int {
		if left > right {
			return -1
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				return binarySearch(nums, left, mid-1, target)
			} else {
				return binarySearch(nums, mid+1, right, target)
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				return binarySearch(nums, mid+1, right, target)
			} else {
				return binarySearch(nums, left, mid-1, target)
			}
		}
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}
