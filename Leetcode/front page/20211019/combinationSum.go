package _0211019

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	temp:=[]int{}
	res:=[][]int{}
	var dfs func(index int,sum int)
	dfs=func(index int,sum int){
		if sum==target{
			t:=make([]int,len(temp))
			copy(t,temp)
			res=append(res,t)
		}
		if sum>target{
			return
		}
		for i:=index;i<len(candidates);i++{
			temp=append(temp,candidates[i])
			dfs(i,sum+candidates[i])
			temp=temp[:len(temp)-1]
		}
	}
	dfs(0,0)
	return res
}
