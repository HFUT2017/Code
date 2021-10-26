package _0211027

func calculate(s string) int {
	stack := []int{}
	num := 0
	preOperator := '+'
	res := 0
	for index, ch := range s {
		isDigist := '0' <= ch && '9' >= ch
		if isDigist {
			num = num*10 + int(ch-'0')
		}
		if !isDigist && ch != ' ' || index == len(s)-1 {
			switch preOperator {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -1*num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			preOperator = ch
			num = 0
		}
	}
	for _, value := range stack {
		res += value
	}
	return res
}
