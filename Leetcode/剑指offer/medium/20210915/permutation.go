package _0210915

import "sort"

// Permutation 输入一个字符串，打印出该字符串中字符的所有排列。
func Permutation(s string) []string {
	res := []string{}
	p := make([]byte, 0, len(s))
	sSlice := []byte(s)
	used := make([]bool, len(s))
	sort.Slice(sSlice, func(i, j int) bool {
		return sSlice[i] < sSlice[j]
	})
	var backTrack func()
	backTrack = func() {
		if len(p) == len(s) {
			res = append(res, string(p))
			return
		}
		for i := 0; i < len(s); i++ {
			if used[i] == true || (i > 0 && used[i-1] == false && sSlice[i] == sSlice[i-1]) {
				continue
			}
			used[i] = true
			p = append(p, sSlice[i])
			backTrack()
			used[i] = false
			p = p[:len(p)-1]
		}
	}
	backTrack()
	return res
}
