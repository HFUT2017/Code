package _0211028

import "strings"

func longestSubstring(s string, k int) int {
	var hash [26]int
	for _, ch := range s {
		hash[ch-'a']++
	}
	for index, value := range hash {
		if 0 < value && value < k {
			res := 0
			for _, val := range strings.Split(s, string('a'+byte(index))) {
				res = max(res, longestSubstring(val, k))
			}
			return res
		}
	}
	return len(s)
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
