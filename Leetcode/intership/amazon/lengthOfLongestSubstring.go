package amazon

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	hash := map[byte]bool{}
	count := 0
	for right < len(s) {
		for left < right && hash[s[right]] == true {
			hash[s[left]] = false
			left++
		}
		hash[s[right]] = true
		if right-left+1 > count {
			count = right - left + 1
		}
		right++
	}
	return count
}
