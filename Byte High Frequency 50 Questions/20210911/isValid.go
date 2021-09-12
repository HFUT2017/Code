package _0210911

func isValid(s string) bool {
	dp := make([]int32, len(s))
	top := 0
	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {
			dp[top] = value
			top++
		} else {
			if top == 0 {
				return false
			}
			top--
			if value == ')' && dp[top] != '(' {
				return false
			}
			if value == ']' && dp[top] != '[' {
				return false
			}
			if value == '}' && dp[top] != '{' {
				return false
			}
		}
	}
	if top != 0 {
		return false
	}
	return true
}
