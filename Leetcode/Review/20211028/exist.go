package _0211028

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < len(used); i++ {
		used[i] = make([]bool, len(board[i]))
	}
	var dfs func(row, col, index int) bool
	dfs = func(row, col, index int) bool {
		if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || used[row][col] == true || board[row][col] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		used[row][col] = true
		defer func() {
			used[row][col] = false
		}()
		for _, value := range dir {
			i, j := row+value[0], col+value[1]
			if dfs(i, j, index+1) == true {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0) == true {
				return true
			}
		}
	}
	return false
}
