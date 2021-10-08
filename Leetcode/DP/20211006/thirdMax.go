package _0211006

func thirdMax(nums []int) int {
	k := 3
	length := len(nums) - 1
	stack := []int{}
	buildHeap(nums)
	if len(nums) < 3 {
		return nums[0]
	}
	for len(stack) != k && length >= 0 {
		if len(stack) == 0 || nums[0] != stack[len(stack)-1] {
			stack = append(stack, nums[0])
		}
		nums[0], nums[length] = nums[length], nums[0]
		length--
		adjustHeap(nums, 0, length)
	}
	if len(stack) < 3 {
		return stack[0]
	} else {
		return stack[len(stack)-1]
	}
}

func buildHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		adjustHeap(nums, i, len(nums)-1)
	}
}

func adjustHeap(nums []int, index int, length int) {
	left, right, maxIndex := 2*index, 2*index+1, index
	if left <= length && nums[left] > nums[maxIndex] {
		maxIndex = left
	}
	if right <= length && nums[right] > nums[maxIndex] {
		maxIndex = right
	}
	if maxIndex != index {
		nums[index], nums[maxIndex] = nums[maxIndex], nums[index]
		adjustHeap(nums, maxIndex, length)
	}
}
