package _0211001

func numlslands(grid [][]byte) int {
	var dfs func(i, j int)
	res := 0
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i, j-1)
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
