package _0211004

func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return isSub(root.Left,root.Right)
}

func isSub(a,b *TreeNode) bool{
	if a==nil&&b==nil{
		return true
	}
	if a==nil||b==nil{
		return false
	}
	if a.Val!=b.Val{
		return false
	}
	return isSub(a.Left,b.Right)&&isSub(a.Right,b.Left)
}
