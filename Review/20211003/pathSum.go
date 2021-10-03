package _0211003

func pathSum(root *TreeNode, target int) [][]int {
	res:=[][]int{}
	var dfs func(root *TreeNode ,temp []int,target int)
	dfs=func(root *TreeNode,temp []int,target int){
		if root==nil{
			return
		}
		target-=root.Val
		temp=append(temp,root.Val)
		if target==0&&root.Left==nil&&root.Right==nil{
			t:=make([]int,len(temp))
			copy(t,temp)
			res=append(res,t)
		}
		dfs(root.Left,temp,target)
		dfs(root.Right,temp,target)
		temp=temp[:len(temp)-1]
	}
	dfs(root,[]int{},target)
	return res
}
