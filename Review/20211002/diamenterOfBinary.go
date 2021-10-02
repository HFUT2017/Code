package _0211002

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	return max(left, right) + 1
}
