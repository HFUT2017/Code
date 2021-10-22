package _0211022

import "sort"

func subsetsWithDup(nums []int) [][]int {
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func(temp []int, index int)
	res := [][]int{}
	dfs = func(temp []int, index int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := index; i < len(nums); i++ {
			if used[i] == true || (i > 0 && used[i-1] == false && nums[i] == nums[i-1]) {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			dfs(temp, i+1)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
