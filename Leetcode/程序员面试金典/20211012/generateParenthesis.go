package _0211012

func generateParenthesis(n int) []string {
	res := []string{}
	var backTrack func(left, right int, temp string)
	backTrack = func(left, right int, temp string) {
		if len(temp) == 2*n {
			res = append(res, temp)
		}
		if left < n {
			temp += "("
			backTrack(left+1, right, temp)
			temp = temp[:len(temp)-1]
		}
		if right < left {
			temp += ")"
			backTrack(left, right+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backTrack(0, 0, "")
	return res
}
