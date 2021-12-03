package intership

func CountGoodSubstrings(s string) int {
	right := 3
	res := 0
	for right <= len(s) {
		str := s[right-3 : right]
		if str[0] != str[1] && str[0] != str[2] && str[1] != str[2] {
			res++
		}
		right++
	}
	return res
}
