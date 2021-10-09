package _0211009

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	res := 0
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m)
		for j := 0; j < len(dp[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				res = 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
