package _0211005

import "math"

func MinFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	for j := 0; j < len(matrix[0]); j++ {
		dp[0][j] = matrix[0][j]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				if j+1 < len(matrix[i]) {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + matrix[i][j]
				}
			} else if j == len(matrix[i])-1 {
				if j-1 >= 0 {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
				} else {
					dp[i][j] = dp[i-1][j] + matrix[i][j]
				}
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i-1][j+1])) + matrix[i][j]
			}
		}
	}
	res := math.MaxInt32
	for j := 0; j < len(matrix[0]); j++ {
		if dp[len(matrix)-1][j] < res {
			res = dp[len(matrix)-1][j]
		}
	}
	return res
}

