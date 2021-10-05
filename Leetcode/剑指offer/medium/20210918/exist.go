package _0210918

func Exist(board [][]byte, word string) bool {
	return helper(board, word, 0, 0, 0)
}

func helper(board [][]byte, word string, i, j, index int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || index >= len(word) {
		return false
	}
	if board[i][j] == word[index] {
		if index == len(word)-1 {
			return true
		}
		res := helper(board, word, i+1, j, index+1)
		if !res {
			res = helper(board, word, i, j+1, index+1)
		}
		if !res {
			res = helper(board, word, i-1, j, index+1)
		}
		if !res {
			res = helper(board, word, i, j-1, index+1)
		}
		return res
	}
	return false
}
