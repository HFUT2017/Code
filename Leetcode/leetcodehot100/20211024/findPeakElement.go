package _0211024

//func findPeakElement(nums []int) int {
//	get := func(index int) int {
//		if index == -1 || index == len(nums) {
//			return math.MinInt32
//		}
//		return nums[index]
//	}
//	left, right := 0, len(nums)-1
//	for left <= right {
//		mid := left + (right-left)/2
//		if get(mid) > get(mid-1) && get(mid) < get(mid+1) {
//			return mid
//		} else if get(mid) < get(mid+1) {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//	return -1
//}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
