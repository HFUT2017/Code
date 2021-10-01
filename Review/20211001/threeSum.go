package _0211001

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -1 * nums[first]
		thrid := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < thrid && nums[second]+nums[thrid] > target {
				thrid--
			}
			if second == thrid {
				break
			}
			if nums[second]+nums[thrid] == target {
				res = append(res, []int{nums[first], nums[second], nums[thrid]})
			}
		}
	}
	return res
}
