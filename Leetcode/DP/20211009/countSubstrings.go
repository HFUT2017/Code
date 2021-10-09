package _0211009

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += center_count(s, i, i)
		res += center_count(s, i, i+1)
	}
	return res
}

func center_count(s string, i, j int) int {
	res := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		res++
		i--
		j++
	}
	return res
}
