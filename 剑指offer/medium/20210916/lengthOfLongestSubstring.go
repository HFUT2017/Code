package _0210916

func LengthOfLongestSubstring(s string) int {
	hash := make(map[byte]bool)
	if len(s) == 0 {
		return 0
	}
	hash[s[0]] = true
	left := 0
	res := 1
	for right := 1; right < len(s); {
		t := hash[s[right]]
		if !t {
			hash[s[right]] = true
			right++
		} else {
			delete(hash, s[left])
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}
