package tsort

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	res := merge(MergeSort(left), MergeSort(right))
	return res
}

func merge(left, right []int) []int {
	res := []int{}
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) || rightIndex < len(right) {
		if leftIndex == len(left) {
			res = append(res, right[rightIndex])
			rightIndex++
			continue
		}
		if rightIndex == len(right) {
			res = append(res, left[leftIndex])
			leftIndex++
			continue
		}
		if left[leftIndex] < right[rightIndex] {
			res = append(res, left[leftIndex])
			leftIndex++
		} else {
			res = append(res, right[rightIndex])
			rightIndex++
		}
	}
	return res
}
