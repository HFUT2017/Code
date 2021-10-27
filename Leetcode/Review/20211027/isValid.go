package _0211027

func isValid(s string) bool {
	stack := []byte{}
	hash := map[byte]byte{']': '[', '}': '{', ')': '('}
	for i := 0; i < len(s); i++ {
		if s[i] == ']' || s[i] == ')' || s[i] == '}' {
			if len(stack) == 0 || hash[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
