package _0211004

func mirrorTree(root *TreeNode) *TreeNode {
	if root==nil{
		return root
	}
	root.Left,root.Right=root.Right,root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
