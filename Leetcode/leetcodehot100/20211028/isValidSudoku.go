package _0211028

func isValidSudoku(board [][]byte) bool {
	var cols, rows [9][9]int
	var subBoxs [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			cols[j][index]++
			subBoxs[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index] > 1 || subBoxs[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
