package _0211017

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var str strings.Builder
	str.WriteString(s2)
	str.WriteString(s2)
	return strings.Contains(str.String(), s1)
}
