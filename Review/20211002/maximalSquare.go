package _0211002

func maximalSquare(matrix [][]byte) int {
	maxSlide := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				maxSlide=1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
				if dp[i][j] > maxSlide {
					maxSlide = dp[i][j]
				}
			}
		}
	}
	return maxSlide * maxSlide

}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
