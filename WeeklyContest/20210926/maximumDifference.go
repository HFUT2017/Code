package _0210926

import "math"

func maximumDifference(nums []int) int {
	res := math.MinInt32
	min := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i+1]-min > res {
			res = nums[i+1] - min
		}
	}
	if res > 0 {
		return res
	}
	return -1
}
