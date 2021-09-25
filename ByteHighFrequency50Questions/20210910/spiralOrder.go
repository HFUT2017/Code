package _0210910

// SpiralOrder 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	n := len(matrix)
	if n == 0 {
		return res
	}
	m := len(matrix[0])
	left, right, up, dowm := 0, m-1, 0, n-1
	for left <= right && up <= dowm {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		for i := up + 1; i <= dowm; i++ {
			res = append(res, matrix[i][right])
		}
		if left < right && up < dowm {
			for i := right - 1; i > left; i-- {
				res = append(res, matrix[dowm][i])
			}
			for i := dowm; i > up; i-- {
				res = append(res, matrix[i][left])
			}
		}
		left++
		right--
		up++
		dowm--
	}
	return res
}
