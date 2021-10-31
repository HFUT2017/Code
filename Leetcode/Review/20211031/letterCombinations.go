package _0211031

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	str := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var dfs func(index int, temp string)
	res := []string{}
	dfs = func(index int, temp string) {
		if len(temp) == len(digits) {
			res = append(res, temp)
			return
		}
		for i := index; i < len(digits); i++ {
			for j := 0; j < len(str[(digits[i]-'0')]); j++ {
				temp += string(str[(digits[i] - '0')][j])
				dfs(i+1, temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(0, "")
	return res
}
