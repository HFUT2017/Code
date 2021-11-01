package _0211101

func isValidSudoku(board [][]byte) bool {
	var cols, rows [9][9]int
	var subBoxs [3][3][9]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				dig := int(board[i][j] - '1')
				rows[i][dig]++
				cols[j][dig]++
				subBoxs[i/3][j/3][dig]++
				if rows[i][dig] > 1 || cols[j][dig] > 1 || subBoxs[i/3][j/3][dig] > 1 {
					return false
				}
			}
		}
	}
	return true
}
