package _0211012

func isUnique(astr string) bool {
	if len(astr) > 26 {
		return false
	}
	hash := map[rune]bool{}
	for _, value := range astr {
		if _, ok := hash[value]; ok {
			return false
		}
		hash[value] = true
	}
	return true
}
