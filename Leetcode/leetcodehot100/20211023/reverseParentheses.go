package _0211023

func reverseParentheses(s string) string {
	str := []byte{}
	for i, _ := range s {
		if s[i] == ')' {
			subStr := []byte{}
			exit := false
			for exit == false {
				c := str[len(str)-1]
				str = str[:len(str)-1]
				if c == '(' {
					exit = true
					str = append(str, subStr...)
					continue
				}
				subStr = append(subStr, c)
			}
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
