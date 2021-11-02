package _0211102

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	ans := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				ans = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
