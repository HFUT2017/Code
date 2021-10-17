package _0211017

import "strconv"

func calculate(s string) int {
	ans, last, start, end := 0, 0, -1, -1
	symbol := byte('+')
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			if end == -1 {
				end = i
			}
			num, _ := strconv.Atoi(s[start:end])
			switch symbol {
			case '+':
				ans += last
				last = num
			case '-':
				ans += last
				last = -num
			case '*':
				last *= num
			case '/':
				last /= num
			}
			start, end = -1, -1
			if i != len(s) {
				symbol = s[i]
			}
		} else if s[i] != ' ' && start == -1 {
			start = i
		} else if s[i] == ' ' && start != -1 && end == -1 {
			end = i
		}
	}
	return ans + last
}
