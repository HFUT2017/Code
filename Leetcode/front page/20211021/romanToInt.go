package _0211021

func romanToInt(s string) int {
	res := 0
	var hash = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := len(s) - 1; i >= 0; i-- {
		if i == len(s)-1 {
			res += hash[s[i]]
			continue
		}
		if hash[s[i]] < hash[s[i+1]] {
			res -= hash[s[i]]
		} else {
			res += hash[s[i]]
		}
	}
	return res
}
