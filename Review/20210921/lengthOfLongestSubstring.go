package Review

//100% 94%
func lengthOfLongestSubstring_uint8(s string) int {
	left, right, res := 0, 0, 0
	var bitset [256]uint8
	for right < len(s) {
		if right < len(s) && bitset[s[right]] == 0 {
			bitset[s[right]] = 1
			right++
		} else {
			bitset[s[left]] = 0
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

//92% 100%
func lengthOfLongestSubstring_byte(s string) int {
	left, right, res := 0, 0, 0
	var hash [256]byte
	for right < len(s) {
		if right < len(s) && hash[s[right]-'a'] == 0 {
			hash[s[right]-'a'] = 1
			right++
		} else {
			hash[s[left]-'a'] = 0
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

//46% 72%
func lengthOfLongestSubstring_map(s string) int {
	left, right, res := 0, 0, 0
	hash := map[byte]int{}
	for right < len(s) {
		if _, ok := hash[s[right]]; !ok {
			hash[s[right]] = 1
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
