package _0211024

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var heights = make([]int, 0)
	result := 0
	for j := 0; j < len(matrix[0]); j++ {
		heights = append(heights, 0)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		result = max(result, largestRectangleArea(heights))
	}
	return result
}
