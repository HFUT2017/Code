package _0210918

// PathSum 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
func PathSum(root *TreeNode, target int) [][]int {
	path:=[]int{}
	res:=[][]int{}
	var dfs func(root *TreeNode,target int)
	dfs=func(root *TreeNode,target int){
		if root==nil{
			return
		}
		target-=root.Val
		path=append(path,root.Val)
		defer func(){path=path[:len(path)-1]}()
		if root.Left==nil&&root.Right==nil&&target==0{
			res=append(res,append([]int(nil),path...))
			return
		}
		dfs(root.Left,target)
		dfs(root.Right,target)
	}
	dfs(root,target)
	return res
}
