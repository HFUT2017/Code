package _0211027

func lengthOfLongestSubstring(s string) int {
	hash := [256]bool{}
	left, right := 0, 0
	res := 0
	for right < len(s) {
		for left < right && hash[s[right]] == true {
			hash[s[left]] = false
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
		hash[s[right]] = true
		right++
	}
	return res
}
