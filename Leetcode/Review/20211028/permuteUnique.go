package _0211028

import "sort"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var backTrack func(temp []int)
	backTrack = func(temp []int) {
		if len(temp) == len(nums) {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true || (i > 0 && used[i-1] == false && nums[i] == nums[i-1]) {
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
