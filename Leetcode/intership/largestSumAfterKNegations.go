package intership

import (
	"math"
	"sort"
)

func LargestSumAfterKNegations(nums []int, k int) int {
	negNum := 0
	minValue := math.MaxInt32
	sum := 0
	sort.Ints(nums)
	for index, value := range nums {
		if value > 0 {
			if index > 0 {
				minValue = min(minValue, min(abs(nums[index-1]), abs(value)))
			} else {
				minValue = min(minValue, abs(value))
			}
			break
		}
		negNum++
	}
	minValue = min(minValue, abs(nums[len(nums)-1]))
	if k > negNum {
		for i := 0; i < negNum; i++ {
			sum += -1 * nums[i]
		}
		for i := negNum; i < len(nums); i++ {
			sum += nums[i]
		}
		println(sum)
		println(minValue)
		if (k-negNum)%2 == 1 {
			sum = sum - 2*abs(minValue)
		}
	} else {
		for i := 0; i < k; i++ {
			sum += -1 * nums[i]
		}
		for i := k; i < len(nums); i++ {
			sum += nums[i]
		}
	}
	return sum
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
