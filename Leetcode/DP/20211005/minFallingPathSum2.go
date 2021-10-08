package _0211005

import "math"

func MinFallingPathSum2(grid [][]int) int {
	n := len(grid[0]) - 1
	if n == 0 {
		return grid[0][0]
	}

	res := math.MaxInt32
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			temp := math.MaxInt32
			for k := 0; k <= n; k++ {
				if k == j {
					continue
				}
				temp = min(temp, grid[i-1][k])
			}
			grid[i][j] += temp
			if i == n {
				res = min(res, grid[i][j])
			}
		}
	}
	return res
}
