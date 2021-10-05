package _0211002

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	res:=0
	dfs=func(root *TreeNode) int{
		if root==nil{
			return 0
		}
		left:=dfs(root.Left)
		right:=dfs(root.Right)
		res=max(res,left+right+1)
		return max(left,right)+1
	}
	dfs(root)
	return res-1
}
