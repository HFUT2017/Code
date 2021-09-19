package _0210919

func ValidateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, value := range pushed {
		stack = append(stack, value)
		for len(stack) > 0 && popped[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return !(len(stack) > 0)
}
