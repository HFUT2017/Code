package _0210912

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	var backTrack func(s string, left, right, n int)
	backTrack = func(s string, left, right, n int) {
		if len(s) == n*2 {
			res = append(res, s)
			return
		}
		if left < n {
			s += "("
			backTrack(s, left+1, right, n)
			s = s[:len(s)-1]
		}
		if right<left {
			s += ")"
			backTrack(s, left, right+1, n)
			s = s[:len(s)-1]
		}
	}
	backTrack("", 0, 0, n)
	return res
}
