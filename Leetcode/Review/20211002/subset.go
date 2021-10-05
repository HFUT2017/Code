package _0211002

func subset(nums []int) [][]int {
	res:=[][]int{}
	var dfs func(temp []int,index int)
	dfs=func(temp []int,index int){
		t:=make([]int,len(temp))
		copy(t,temp)
		res=append(res,t)
		for i:=index;i<len(nums);i++{
			temp=append(temp,nums[i])
			dfs(temp,i+1)
			temp=temp[:len(temp)-1]
		}
	}
	dfs([]int{},0)
	return res
}
