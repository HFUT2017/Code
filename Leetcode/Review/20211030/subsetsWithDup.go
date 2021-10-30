package _0211030

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var backTrack func(index int, temp []int)
	backTrack = func(index int, temp []int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := index; i < len(nums); i++ {
			if i > 0 && used[i-1] == false && nums[i] == nums[i-1] {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			backTrack(i+1, temp)
			temp = temp[:len(temp)-1]
			used[i] = false
		}
	}
	backTrack(0, []int{})
	return res
}
