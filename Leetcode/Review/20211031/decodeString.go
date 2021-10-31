package _0211031

import "strings"

func decodeString(s string) string {
	strStack := []string{}
	numStack := []int{}
	count := 0
	res := ""
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			count = count*10 + int(ch-'0')
		} else if ch == '[' {
			strStack = append(strStack, res)
			numStack = append(numStack, count)
			count = 0
			res = ""
		} else if ch == ']' {
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			s := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			res = s + strings.Repeat(res, repeat)
		} else {
			res += string(ch)
		}
	}
	return res
}
