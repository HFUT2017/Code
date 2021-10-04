package _0211004

func firstUniqChar(s string) byte {
	hash:=map[byte]int{}
	for i:=0;i<len(s);i++{
		hash[s[i]]++
	}
	for i:=0;i<len(s);i++{
		if hash[s[i]]==1{
			return s[i]
		}
	}
	return ' '
}
