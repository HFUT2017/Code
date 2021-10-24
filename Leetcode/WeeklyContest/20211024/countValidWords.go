package _0211024

import "strings"

func countValidWords(sentence string) int {
	str := strings.Split(sentence, " ")
	res := 0
	check := func(str string) bool {
		if str==""||str[0]==' ' {
			return false
		}
		cnt := 0
		for i := 0; i < len(str); i++ {
			if str[i] == '-' {
				cnt++
				if cnt > 1 {
					return false
				}
				if i == 0 || i == len(str)-1 {
					return false
				}
				if str[i-1] < 'a' || str[i-1] > 'z' || str[i+1] < 'a' || str[i+1] > 'z' {
					return false
				}
			} else if str[i] >= '0' && str[i] <= '9' {
				return false
			} else if str[i] < 'a' || str[i] > 'z' {
				if i != len(str)-1 {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < len(str); i++ {
		if check(str[i]) {
			res++
		}
	}
	return res
}
