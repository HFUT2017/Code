package _0210930

func lengthOfLongestSubstring(s string) int {
	hash := map[byte]bool{}
	left, right, res := 0, 0, 0
	for right < len(s) {
		for left < right && hash[s[right]] == true {
			delete(hash, s[left])
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
