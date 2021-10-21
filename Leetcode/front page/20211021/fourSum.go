package _0211021

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			t := target - (nums[first] + nums[second])
			four := len(nums) - 1
			for third := second + 1; third < len(nums); third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				for third < four && nums[third]+nums[four] > t {
					four--
				}
				if third == four {
					break
				}
				if nums[third]+nums[four] == t {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
		}
	}
	return res
}
