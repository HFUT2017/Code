package _0211023

import "strings"

func decodeString(str string) string {
	numStack := []int{}
	strStack := []string{}
	count := 0
	res := ""
	for i := 0; i < len(str); i++ {
		n := str[i] - '0'
		if n >= 0 && n <= 9 {
			count = count*10 + int(n)
		} else if str[i] == '[' {
			numStack = append(numStack, count)
			strStack = append(strStack, res)
			count = 0
			res = ""
		} else if str[i] == ']' {
			numRepeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			s := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			res = s + strings.Repeat(res, numRepeat)
		} else {
			res += string(str[i])
		}
	}
	return res
}
