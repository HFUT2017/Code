package _0210912

func reversePrefix(word string, ch byte) string {
	i := 0
	for ; i < len(word); i++ {
		if word[i] == ch {
			break
		}
	}
	if i == len(word) {
		return word
	}
	str := reverse(word[:i+1])
	return str + word[i+1:]

}

func reverse(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}
