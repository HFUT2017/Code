package _0210927

func findKthLargest(nums []int, k int) int {
	buildHeap(nums)
	length := len(nums)
	for k != 0 {
		k--
		if k == 0 {
			return nums[0]
		}
		nums[0], nums[length-1] = nums[length-1], nums[0]
		length--
		adjustHeap(nums, length, 0)
	}
	return 0
}

func buildHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		adjustHeap(nums, len(nums), i)
	}
}

func adjustHeap(nums []int, length, index int) {
	left, right, max := index*2, index*2+1, index
	if left < length && nums[left] > nums[max] {
		max = left
	}
	if right < length && nums[right] > nums[max] {
		max = right
	}
	if max != index {
		nums[max], nums[index] = nums[index], nums[max]
		adjustHeap(nums, length, max)
	}
}
