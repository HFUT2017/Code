package _0211024

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	heights = append(heights, 0)
	res := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			res = max(res, (i-l-1)*heights[top])
		}
		stack = append(stack, i)

	}
	return res
}
