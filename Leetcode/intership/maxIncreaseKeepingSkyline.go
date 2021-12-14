package intership

func MaxIncreaseKeepingSkyline(grid [][]int) int {
	row, col := []int{}, []int{}
	increasedNUm := 0
	for i := 0; i < len(grid); i++ {
		row = append(row, grid[i][0])
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				col = append(col, grid[i][j])
			} else {
				if grid[i][j] > col[j] {
					col[j] = grid[i][j]
				}
			}
			if grid[i][j] > row[i] {
				row[i] = grid[i][j]
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			increasedNUm += min(row[i], col[j]) - grid[i][j]
		}
	}
	return increasedNUm
}
