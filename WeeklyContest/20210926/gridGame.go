package _0210926

var (
	used1 [][]bool
	max   int
)

func GridGame(grid [][]int) int64 {
	max = 0
	used := make([][]bool, 2)
	used[0] = make([]bool, len(grid[0]))
	used[1] = make([]bool, len(grid[0]))
	used1 = make([][]bool, 2)
	used1[0] = make([]bool, len(grid[0]))
	used1[1] = make([]bool, len(grid[0]))
	dfs(grid, 0, 0, 0, used)
	for i := 0; i < len(used1); i++ {
		for j := 0; j < len(used1[i]); j++ {
			if used1[i][j] == true {
				grid[i][j] = 0
				used1[i][j] = false
			}
		}
	}
	max = 0
	dfs(grid, 0, 0, 0, used1)
	return int64(max)
}

func dfs(grid [][]int, i, j int, res int, used [][]bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	res += grid[i][j]
	used[i][j] = true
	if i == 1 && j == len(grid[0])-1 {
		if res > max {
			for k := 0; k < len(used[0]); k++ {
				used1[0][k] = used[0][k]
				used1[1][k] = used[1][k]
			}
			max = res
		}
	}
	dfs(grid, i+1, j, res, used)
	dfs(grid, i, j+1, res, used)
	res -= grid[i][j]
	used[i][j] = false
}
