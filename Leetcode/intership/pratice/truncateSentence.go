package intership

import "strings"

func truncateSentence(s string, k int) string {
	str := strings.Split(s, " ")
	res := ""
	for i := 0; i < k; i++ {
		res += str[i]
		if i != k-1 {
			res += " "
		}
	}
	return res
}
