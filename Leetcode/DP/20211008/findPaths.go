package _0211008

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	mem := make([][][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([][]int, n)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = make([]int, maxMove+1)
		}
	}
	for i := 0; i < len(mem); i++ {
		for j := 0; j < len(mem[i]); j++ {
			for k := 0; k < len(mem[i][j]); k++ {
				mem[i][j][k] = -1
			}
		}
	}
	var dfs func(i, j int, maxMove int) int
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(i, j int, maxMove int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 1
		}
		if maxMove == 0 {
			return 0
		}
		if mem[i][j][maxMove] != -1 {
			return mem[i][j][maxMove]
		}
		ans := 0
		for index := 0; index < len(dir); index++ {
			ans += dfs(i+dir[index][0], j+dir[index][1], maxMove-1)
			ans = ans % 1000000007
		}
		mem[i][j][maxMove] = ans
		return mem[i][j][maxMove]
	}
	return dfs(startRow, startColumn, maxMove)
}
