package _0211021

func searchInsert(nums []int, target int) int {
	var binarySearch func(nums []int, target int) int
	binarySearch = func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (left + right) >> 1
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	return binarySearch(nums, target)
}
