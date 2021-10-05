package _0210924

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])
	res := []int{}
	left, right, up, down := 0, m-1, 0, n-1
	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}

		for i := up + 1; i <= down; i++ {
			res = append(res, matrix[i][right])
		}

		if left < right && up < down {
			for i := right - 1; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			for i := down - 1; i > up; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left, right, up, down = left+1, right-1, up+1, down-1
	}
	return res
}
