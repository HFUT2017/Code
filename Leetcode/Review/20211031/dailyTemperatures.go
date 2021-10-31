package _0211031

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			res[preIndex] = i - preIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
