package _0211102

import "strconv"

func restoreIpAddresses(s string) []string {
	segments := make([]int, 4)
	res := []string{}
	var dfs func(index int, start int)
	dfs = func(index int, start int) {
		if index == 4 {
			if start == len(s) {
				ip := ""
				for i := 0; i < len(segments); i++ {
					ip += strconv.Itoa(segments[i])
					if i != len(segments)-1 {
						ip += "."
					}
				}
				res = append(res, ip)
			}
			return
		}
		if start == len(s) {
			return
		}
		if s[start] == '0' {
			segments[index] = 0
			dfs(index+1, start+1)
		}
		addr := 0
		for end := start; end < len(s); end++ {
			addr = addr*10 + int(s[end]-'0')
			if addr > 0 && addr <= 255 {
				segments[index] = addr
				dfs(index+1, end+1)
			} else {
				break
			}
		}
	}
	dfs(0,0)
	return res
}
