package _0211025

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Right)
		sum += root.Val
		root.Val = sum
		inorder(root.Left)
	}
	inorder(root)
	return root
}
