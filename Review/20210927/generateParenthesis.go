package _0210927

func generateParenthesis(n int) []string {
	res := []string{}
	var backTrack func(left, right int, temp string)
	backTrack = func(left, right int, temp string) {
		if len(temp) == 2*n {
			res = append(res, temp)
			return
		}
		if left < n {
			temp = temp + "("
			backTrack(left+1, right, temp)
			temp = temp[:len(temp)-1]
		}
		if right < left {
			temp = temp + ")"
			backTrack(left, right+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backTrack(0, 0, "")
	return res

}
