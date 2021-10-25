package _0211026

func findAnagrams(s string, p string) []int {
	hash:=[26]int{}
	count:=[26]int{}
	res:=[]int{}
	for i:=0;i<len(p);i++{
		hash[p[i]-'a']++
	}
	for left,right:=0,0;right<len(s);right++{
		count[s[right]-'a']++
		for count[s[right]-'a']>hash[s[right]-'a']{
			count[s[left]-'a']--
			left++
		}
		if right-left+1==len(p){
			res=append(res,left)
		}
	}
	return res
}
