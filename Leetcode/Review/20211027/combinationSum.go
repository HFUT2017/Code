package _0211027

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var dfs func(index int, temp []int, sum int)
	dfs = func(index int, temp []int, sum int) {
		if sum == target {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			dfs(i+1, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}
