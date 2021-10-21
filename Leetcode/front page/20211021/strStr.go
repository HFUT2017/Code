package _0211021

import "strings"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	res := strings.Split(haystack, needle)
	if len(res[0]) == len(haystack) {
		return -1
	} else {
		return len(res[0])
	}
}
