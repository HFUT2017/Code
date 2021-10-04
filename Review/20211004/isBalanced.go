package _0211004

func getDepth(root *TreeNode) int{
	if root==nil{
		return 0
	}
	left:=getDepth(root.Left)
	right:=getDepth(root.Right)
	return max(left,right)+1
}
func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	left:=getDepth(root.Left)
	right:=getDepth(root.Right)
	if left<right{
		left,right=right,left
	}
	return left-right<=1&&isBalanced(root.Left)&&isBalanced(root.Right)
}
