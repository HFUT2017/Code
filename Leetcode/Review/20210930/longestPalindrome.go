package _0210930

func longestPalindrome(s string) string {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := center(s, i, i)
		left2, right2 := center(s, i, i+1)
		if right1-left1 > right-left {
			right = right1
			left = left1
		}
		if right2-left2 > right-left {
			right = right2
			left = left2
		}
	}
	return s[left : right+1]
}

func center(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - 1
}
