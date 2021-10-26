package _0211028

func solve(board [][]byte) {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < len(board); i++ {
		dfs(i, 0)
		dfs(i, len(board[i])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(0, j)
		dfs(len(board)-1, j)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
