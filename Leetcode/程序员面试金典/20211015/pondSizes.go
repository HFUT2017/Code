package _0211015

import "sort"

func pondSizes(land [][]int) []int {
	var dfs func(i, j int) int
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= len(land) || j >= len(land[0]) || land[i][j] != 0 {
			return 0
		}
		land[i][j] = 1
		res := 1
		for k := 0; k < len(dir); k++ {
			res += dfs(i+dir[k][0], j+dir[k][1])
		}
		return res
	}
	res := []int{}
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			ans := dfs(i, j)
			if ans != 0 {
				res = append(res, ans)
			}
		}
	}
	sort.Ints(res)
	return res
}
