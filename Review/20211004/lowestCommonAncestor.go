package _0211004

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==p||root==q||root==nil{
		return root
	}
	left:=lowestCommonAncestor(root.Left,p,q)
	right:=lowestCommonAncestor(root.Right,p,q)
	if left==nil{
		return right
	}
	if right==nil{
		return left
	}
	return root
}