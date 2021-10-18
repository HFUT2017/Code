package _0211019

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	temp:=[]int{}
	res:=[][]int{}
	used:=make([]bool,len(candidates))
	var dfs func(index int,sum int)
	dfs=func(index int,sum int){
		if sum==target{
			t:=make([]int,len(temp))
			copy(t,temp)
			res=append(res,t)
			return
		}
		if sum>target{
			return
		}
		for i:=index;i<len(candidates);i++{
			if used[i]==true{
				continue
			}
			if i!=0&&candidates[i]==candidates[i-1]&&used[i-1]==false{
				continue
			}
			used[i]=true
			temp=append(temp,candidates[i])
			dfs(i+1,sum+candidates[i])
			temp=temp[:len(temp)-1]
			used[i]=false
		}
	}
	dfs(0,0)
	return res
}
