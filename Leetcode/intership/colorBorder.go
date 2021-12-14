package intership

func ColorBorder(grid [][]int, row int, col int, color int) [][]int {
	target := grid[row][col]
	hash := make([][]bool, len(grid))
	gridbak := make([][]int, len(grid))
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < len(hash); i++ {
		hash[i] = make([]bool, len(grid[i]))
		gridbak[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(gridbak); i++ {
		for j := 0; j < len(gridbak[i]); j++ {
			gridbak[i][j] = grid[i][j]
		}
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || gridbak[i][j] != target {
			return false
		}
		if hash[i][j] == true {
			return true
		}
		hash[i][j] = true
		for _, value := range dir {
			if dfs(i+value[0], j+value[1]) == false {
				grid[i][j] = color
			}
		}
		return true
	}
	dfs(row, col)
	return grid
}
