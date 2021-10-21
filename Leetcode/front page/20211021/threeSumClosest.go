package _0211021

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		second := first + 1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				third--
			} else {
				second++
			}

		}
	}
	return res
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -1 * num
}
