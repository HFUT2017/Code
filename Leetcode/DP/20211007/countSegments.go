package _0211007

func countSegments(s string) int {
	res := []string{}
	temp := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			temp += string(s[i])
		} else {
			if temp != "" {
				res = append(res, temp)
			}
			temp = ""
		}
	}
	if temp != "" {
		res = append(res, temp)
	}
	return len(res)
}
