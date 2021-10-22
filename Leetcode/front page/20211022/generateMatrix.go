package _0211022

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	num := 1
	left, right, up, down := 0, n-1, 0, n-1
	for left <= right && up <= down {
		for i := left; i <= right; i++ {
			res[up][i] = num
			num++
		}
		for i := up + 1; i <= down; i++ {
			res[i][right] = num
			num++
		}
		if left < right && up < down {
			for i := right - 1; i >= left; i-- {
				res[down][i] = num
				num++
			}
			for i := down - 1; i > up; i-- {
				res[i][left] = num
				num++
			}
		}
		left++
		right--
		up++
		down--
	}
	return res
}
