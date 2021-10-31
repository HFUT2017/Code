package _0211031

import "strings"

func strStr(haystack string, needle string) int {
	if haystack == "" || needle == "" {
		if needle != "" {
			return -1
		}
		return 0
	}
	str := strings.Split(haystack, needle)
	if len(str[0]) == len(haystack) {
		return -1
	}
	return len(str[0])
}
