package tsort

func HeapSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	buildMaxHeap(nums)
	for length > 0 {
		nums[0], nums[length-1] = nums[length-1], nums[0]
		length--
		adjustHeap(nums, 0, length)
	}
	return
}

func buildMaxHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		adjustHeap(nums, i, len(nums))
	}
}

func adjustHeap(nums []int, index, length int) {
	leftIndex, rightIndex, maxIndex := 2*index, 2*index+1, index
	if leftIndex < length && nums[leftIndex] > nums[maxIndex] {
		maxIndex = leftIndex
	}
	if rightIndex < length && nums[rightIndex] > nums[maxIndex] {
		maxIndex = rightIndex
	}
	if maxIndex != index {
		nums[maxIndex], nums[index] = nums[index], nums[maxIndex]
		adjustHeap(nums, maxIndex, length)
	}
}
