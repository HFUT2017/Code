package _0210929

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxStride := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxStride = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
				if dp[i][j] > maxStride {
					maxStride = dp[i][j]
				}
			}
		}
	}
	return maxStride * maxStride
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
