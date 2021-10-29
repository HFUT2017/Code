package _0211029

func longestIncreasingPath(matrix [][]int) int {
	res := 0
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	mem := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		mem[i] = make([]int, len(matrix[i]))
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if mem[i][j] != 0 {
			return mem[i][j]
		}
		mem[i][j]++
		for k := 0; k < len(dir); k++ {
			newI, newJ := i+dir[k][0], j+dir[k][1]
			if newI < 0 || newJ < 0 || newI >= len(matrix) || newJ >= len(matrix[0]) {
				continue
			}
			if matrix[newI][newJ] <= matrix[i][j] {
				continue
			}
			mem[i][j] = max(mem[i][j], dfs(newI, newJ)+1)
		}
		return mem[i][j]
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}
