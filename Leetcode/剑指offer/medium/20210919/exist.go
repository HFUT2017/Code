package _0210919

func Exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if helper(board, word, i, j, 0, used) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, word string, i, j, index int, used [][]bool) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || index >= len(word) || used[i][j] == true {
		return false
	}
	if board[i][j] == word[index] {
		if index == len(word)-1 {
			return true
		}
		used[i][j] = true
		res := helper(board, word, i+1, j, index+1, used)
		if !res {
			res = helper(board, word, i, j+1, index+1, used)
		}
		if !res {
			res = helper(board, word, i-1, j, index+1, used)
		}
		if !res {
			res = helper(board, word, i, j-1, index+1, used)
		}
		used[i][j] = false
		return res
	}
	return false
}
