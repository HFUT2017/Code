package _0211001

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack)==0{
				return false
			}
			if s[i] == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if s[i] == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if s[i] == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			}else{
				stack=append(stack,s[i])
			}
		}
	}
	return len(stack) == 0
}
