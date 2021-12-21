package amazon

func numIslands(grid [][]byte) int {
	count := 0
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		for _, d := range dir {
			dfs(i+d[0], j+d[1])
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count += 1
				dfs(i, j)
			}
		}
	}
	return count
}
