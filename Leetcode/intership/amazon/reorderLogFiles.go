package amazon

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	letter := []string{}
	digit := []string{}
	res := []string{}
	for _, str := range logs {
		if str[len(str)-1] >= '0' && str[len(str)-1] <= '9' {
			digit = append(digit, str)
		} else {
			letter = append(letter, str)
		}
	}
	sort.Slice(letter, func(i, j int) bool {
		str1, str2 := strings.Split(letter[i], " "), strings.Split(letter[j], " ")
		if letter[i][len(str1[0]):] == letter[j][len(str2[0]):] {
			return str1[0] < str2[0]
		}
		return letter[i][len(str1[0]):] < letter[j][len(str2[0]):]
	})
	for _, let := range letter {
		res = append(res, let)
	}
	for _, dig := range digit {
		res = append(res, dig)
	}
	return res
}
