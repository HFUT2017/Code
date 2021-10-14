package _0211014

func CheckPermutation(s1 string, s2 string) bool {
	hash := map[rune]int{}
	for _, v := range s1 {
		hash[v]++
	}
	for _, v := range s2 {
		hash[v]--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
