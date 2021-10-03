package _0211003

import "sort"

func permutation(s string) []string {
	used:=make([]bool,len(s))
	res:=[]string{}
	str:=[]byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i]<str[j]
	})
	var dfs func(temp string)
	dfs=func(temp string){
		if len(temp)==len(str){
			res=append(res,temp)
			return
		}
		for i:=0;i<len(str);i++{
			if used[i]==true||(i!=0&&used[i-1]==false&&str[i]==str[i-1]){
				continue
			}
			used[i]=true
			temp+=string(str[i])
			dfs(temp)
			used[i]=false
			temp=temp[:len(temp)-1]
		}
	}
	dfs("")
	return res
}
