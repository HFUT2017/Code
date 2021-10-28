package _0211028

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	strs := strings.Split(path, "/")
	for i := 0; i < len(strs); i++ {
		switch strs[i] {
		case ".", "":
			continue
		case "..":
			if len(stack) >= 1 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, strs[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}
