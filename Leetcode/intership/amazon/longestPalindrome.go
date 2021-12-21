package amazon

func longestPalindrome(s string) string {
	str := helper(s)
	radius := make([]int, len(str))
	r, c := -1, -1
	start, maxLen := -1, 0
	for i := 0; i < len(str); i++ {
		if i < r {
			radius[i] = min(radius[2*c-i], r-i)
		} else {
			radius[i] = 1
		}
		for i+radius[i] < len(str) && i-radius[i] > -1 {
			if str[i+radius[i]] == str[i-radius[i]] {
				radius[i]++
			} else {
				break
			}
		}
		if r < i+radius[i] {
			r = i + radius[i]
			c = i
		}
		if maxLen < radius[i] {
			start = (i - radius[i] + 1) / 2
			maxLen = radius[i] - 1
		}
	}
	return s[start : start+maxLen]
}

func helper(s string) string {
	str := make([]byte, 2*len(s)+1)
	for i := 0; i < 2*len(s)+1; i++ {
		if i%2 == 1 {
			str[i] = s[(i-1)/2]
		} else {
			str[i] = '#'
		}
	}
	return string(str)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
