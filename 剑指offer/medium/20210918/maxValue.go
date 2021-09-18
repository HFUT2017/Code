package _0210918

// MaxValue 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
func MaxValue(grid [][]int) int {

	dp := make([][]int, len(grid)+1)
	for i := 0; i <= len(grid); i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i][j-1], dp[i-1][j])
		}
	}

	return dp[len(grid)][len(grid[0])]
}
