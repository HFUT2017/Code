package _0211015

func canPermutePalindrome(s string) bool {
	f := true
	hash := map[rune]int{}
	for _, v := range s {
		hash[v]++
	}
	for _, v := range hash {
		if v%2 == 1 {
			if f == false {
				return false
			}
			f = false
		}
	}
	return true
}
