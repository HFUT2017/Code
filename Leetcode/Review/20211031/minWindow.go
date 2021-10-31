package _0211031

func minWindow(s string, t string) string {
	hash := [256]int{}
	count := [256]int{}
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	check := func() bool {
		for i := 0; i < len(t); i++ {
			if count[t[i]] < hash[t[i]] {
				return false
			}
		}
		return true
	}
	left, right, res := 0, 0, ""
	for right < len(s) {
		count[s[right]]++
		for left < right && check() {
			if res == "" || right-left+1 < len(res) {
				res = s[left : right+1]
			}
			count[s[left]]--
			left++
		}
		right++
	}
	return res
}
