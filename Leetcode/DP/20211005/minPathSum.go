package _0211005

// MinPathSum 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
func MinPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i-1 < 0 && j-1 < 0 {
				dp[i][j] = grid[i][j]
			} else if i-1 < 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
