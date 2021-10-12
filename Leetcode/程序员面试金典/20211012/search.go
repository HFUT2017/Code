package _0211012

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[left] == target {
			return left
		} else if nums[mid] == target {
			right = mid
		} else if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[right] > nums[mid] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			for left <= right && nums[left] == nums[mid] {
				left++
			}
			for left <= right && nums[mid] == nums[right] {
				right--
			}
		}
	}
	return -1
}
