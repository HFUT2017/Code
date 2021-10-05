package _0211003

func lengthOfLongestSubstring(s string) int {
	hash:=map[byte]bool{}
	res:=0
	for left,right:=0,0;right<len(s);right++{
		for hash[s[right]]==true{
			hash[s[left]]=false
			left++
		}
		hash[s[right]]=true
		if right-left+1>res{
			res=right-left+1
		}
	}
	return res
}
