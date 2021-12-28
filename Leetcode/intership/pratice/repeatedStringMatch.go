package intership

import "strings"

func repeatedStringMatch(a string, b string) int {
	m, n, exist := len(a), len(b), make([]bool, 26)
	for _, ch := range a {
		exist[ch-'a'] = true
	}
	for _, ch := range b {
		if !(exist[ch-'a']) {
			return -1
		}
	}
	ret := n / m
	str := strings.Repeat(a, ret)
	for i := 0; i < 3; i++ {
		if strings.Contains(str, b) {
			return ret + i
		}
		str += a
	}
	return -1
}
