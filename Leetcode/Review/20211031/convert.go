package _0211031

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([]string, numRows)
	n := 2*numRows - 2
	for index, value := range s {
		x := index % n
		res[min(x, n-x)] += string(value)
	}
	return strings.Join(res, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
