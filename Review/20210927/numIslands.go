package _0210927

func NumIslands(grid [][]byte) int {
	dir := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	res := 0
	var dfs func(i, j int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		for i := 0; i < len(dir); i++ {
			newx := x + dir[i][0]
			newy := y + dir[i][1]
			dfs(newx, newy)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res += 1
				dfs(i, j)
			}
		}
	}
	return res
}
