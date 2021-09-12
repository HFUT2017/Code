package _0210912

func FindKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeap(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeap(nums, i, heapSize)
	}
}

func maxHeap(nums []int, i, heapSize int) {
	l, r, largest := 2*i+1, 2*i+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeap(nums, largest, heapSize)
	}
}
