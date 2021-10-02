package _0211002

func exit(board [][]byte, word string) bool {
	flag := false
	var backTrack func(i, j int, temp string)
	backTrack = func(i, j int, temp string) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || flag == true {
			return
		}
		temp += string(board[i][j])
		if temp == word {
			flag = true
			return
		}
		backTrack(i+1, j, temp)
		backTrack(i, j+1, temp)
		temp = temp[:len(temp)-1]
	}
	return flag
}
