package intership

import "math"

func minimumDeletions(nums []int) int {
	maxNum, minNum := math.MinInt32, math.MaxInt32
	maxIndex, minIndex := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
		if nums[i] < minNum {
			minNum = nums[i]
			minIndex = i
		}
	}
	if minIndex > maxIndex {
		minIndex, maxIndex = maxIndex, minIndex
	}
	return min(len(nums)-maxIndex+minIndex+1, min(maxIndex+1, len(nums)-minIndex))
}
