package _0211005

import "math"

// MinimumTotal 给定一个三角形 triangle ，找出自顶向下的最小路径和。
func MinimumTotal(triangle [][]int) int {
	dp := [][]int{}
	for i := 0; i < len(triangle); i++ {
		temp := []int{}
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				if i == 0 {
					temp = append(temp, triangle[i][j])
				} else {
					temp = append(temp, triangle[i][j]+dp[i-1][j])
				}
			} else if j == i {
				temp = append(temp, triangle[i][j]+dp[i-1][j-1])
			} else {
				t := min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
				temp = append(temp, t)
			}
		}
		dp = append(dp, temp)
	}
	res := math.MaxInt32
	for i := 0; i < len(dp[len(triangle)-1]); i++ {
		if dp[len(triangle)-1][i] < res {
			res = dp[len(triangle)-1][i]
		}
	}
	return res
}
