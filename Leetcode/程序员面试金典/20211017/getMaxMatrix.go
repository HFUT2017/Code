package _0211017

import "math"

func getMaxMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 4)
	col := make([]int, n)
	maxSum := math.MinInt32
	for i := 0; i < m; i++ {
		col = make([]int, n)
		for j := i; j < m; j++ {
			var sum, startCol int
			for k := 0; k < n; k++ {
				col[k] += matrix[j][k]
				if sum < 0 {
					sum = 0
					startCol = k
				}
				sum += col[k]
				if sum > maxSum {
					maxSum = sum
					res = []int{i, startCol, j, k}
				}
			}
		}
	}
	return res
}
