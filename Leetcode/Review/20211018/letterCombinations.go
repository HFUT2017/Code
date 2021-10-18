package _0211018

func letterCombinations(digits string) []string {
	res := []string{}
	value := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if digits == "" {
		return res
	}
	var dfs func(index int, temp string)
	dfs = func(index int, temp string) {
		if len(temp) == len(digits) {
			res = append(res, temp)
			return
		}
		for i := index; i < len(digits); i++ {
			in := int(digits[i] - '0')
			for j := 0; j < len(value[in]); j++ {
				temp += string(value[in][j])
				dfs(i+1, temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(0, "")
	return res
}
