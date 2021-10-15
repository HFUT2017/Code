package _0211015

func permutation(S string) []string {
	used := make([]bool, len(S))
	res := []string{}
	var backTrack func(temp string)
	backTrack = func(temp string) {
		if len(temp) == len(S) {
			res = append(res, temp)
			return
		}
		for i := 0; i < len(S); i++ {
			if used[i] == true {
				continue
			}
			used[i] = true
			temp += string(S[i])
			backTrack(temp)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	backTrack("")
	return res
}
