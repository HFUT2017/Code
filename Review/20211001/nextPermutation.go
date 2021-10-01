package _0211001

import "sort"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for nums[i] > nums[i+1] {
		i--
	}
	j := len(nums) - 1
	for nums[j] < nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	sort.Ints(nums[i:])
	return
}
