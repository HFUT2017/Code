package _0211012

func smallestK(nums []int, k int) []int {
	res := []int{}
	buildHeap(nums)
	length := len(nums) - 1
	for k != 0 {
		t := nums[0]
		res = append(res, t)
		nums[0], nums[length] = nums[length], nums[0]
		length--
		adjustHeap(nums, 0, length)
		k--
	}
	return res
}

func buildHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		adjustHeap(nums, i, len(nums)-1)
	}
}

func adjustHeap(nums []int, index int, length int) {
	left, right, minIndex := 2*index, 2*index+1, index
	if left <= length && nums[left] < nums[minIndex] {
		minIndex = left
	}
	if right <= length && nums[right] < nums[minIndex] {
		minIndex = right
	}
	if minIndex != index {
		nums[index], nums[minIndex] = nums[minIndex], nums[index]
		adjustHeap(nums, minIndex, length)
	}
}
