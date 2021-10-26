package _0211027

func firstUniqChar(s string) int {
	hash := [256]int{}
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if hash[s[i]] == 1 {
			return i
		}
	}
	return -1
}
