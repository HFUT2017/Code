package _0211015

import "sort"

func permuteUnique(nums []int) [][]int {
	used := make([]bool, len(nums))
	res := [][]int{}
	sort.Ints(nums)
	var backTrack func(temp []int)
	backTrack = func(temp []int) {
		if len(temp) == len(nums) {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true || (i != 0 && nums[i-1] == nums[i] && used[i-1] == false) {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			backTrack(temp)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	backTrack([]int{})
	return res
}
