package Review

//96% 100%
func longestPalindrome(s string) string {
	left, right, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := center(s, i, i)
		left2, right2 := center(s, i, i+1)
		if right1-left1 > res {
			res = right1 - left1
			left = left1
			right = right1

		}
		if right2-left2 > res {
			res = right2 - left2
			left = left2
			right = right2
		}
	}
	return s[left : right+1]
}

func center(s string, i, j int) (int, int) {
	for i >= 0 && i < len(s) && j >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	i++
	j--
	return i, j
}
