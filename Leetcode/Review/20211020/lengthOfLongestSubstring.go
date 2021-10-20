package _0211020

func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	hash := map[byte]bool{}
	for right < len(s) {
		for left < right && hash[s[right]] == true {
			delete(hash, s[left])
			left++
		}
		hash[s[right]] = true
		right++
		if right-left > res {
			res = right - left
		}
	}
	return res
}
