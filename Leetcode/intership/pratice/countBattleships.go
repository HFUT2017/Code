package intership

func countBattleships(board [][]byte) int {
	count := 0
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] != 'X' {
			return
		}
		board[i][j] = '.'
		for _, d := range dir {
			dfs(i+d[0], j+d[1])
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}
