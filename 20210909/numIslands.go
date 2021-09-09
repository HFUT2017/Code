package _0210909

var direction = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

// NumIslands 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
func NumIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		grid[row][col] = '0'
		for i := 0; i < 4; i++ {
			nextRow, nextCol := row + direction[i][0], col + direction[i][1]
			if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || grid[nextRow][nextCol] == '0' {
				continue
			}
			dfs(nextRow, nextCol)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
