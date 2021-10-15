package _0211015

func setZeroes(matrix [][]int) {
	dp := make([][]bool, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = true
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] {
				for row := 0; row < len(matrix); row++ {
					matrix[row][j] = 0
				}
				for col := 0; col < len(matrix[i]); col++ {
					matrix[i][col] = 0
				}
			}
		}
	}
}
