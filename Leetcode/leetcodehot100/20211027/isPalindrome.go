package _0211027

func isPalindrome(s string) bool {
	str := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			str = append(str, s[i]-'A'+'a')
		} else if s[i] >= 'a' && s[i] <= 'z' {
			str = append(str, s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			str = append(str, s[i])
		}
	}
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
