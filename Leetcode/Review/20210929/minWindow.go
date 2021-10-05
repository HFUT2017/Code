package _0210929

func minWindow(s string, t string) string {
	hash := [256]int{}
	count := [256]int{}
	res := ""
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	check := func() bool {
		for index, value := range hash {
			if count[index] < value {
				return false
			}
		}
		return true
	}
	for left, right := 0, 0; right < len(s); right++ {
		count[s[right]]++
		for check() && left <= right {
			if res == "" || right-left+1 < len(res) {
				res = s[left : right+1]
			}
			count[s[left]]--
			left++
		}
	}
	return res
}
