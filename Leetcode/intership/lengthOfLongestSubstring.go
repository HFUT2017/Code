package intership

func lengthOfLongestSubstring(s string) int {
	res := 0
	hash := [256]bool{}
	left, right := 0, 0
	for right < len(s) {
		for left < right && hash[s[right]] == true {
			hash[s[left]] = false
			left++
		}
		hash[s[right]] = true
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
	}
	return res
}
