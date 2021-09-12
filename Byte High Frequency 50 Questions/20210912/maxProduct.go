package _0210912

func MaxProduct(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		s1 := longestPalindrome(s[:i])
		s2 := longestPalindrome(s[i+1:])
		res = max(res, len(s1)*len(s2))
	}
	return res
}

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
