package _0211015

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix[0]) - 1
	i, j := len(matrix)-1, 0
	for i >= 0 && j <= m {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}
