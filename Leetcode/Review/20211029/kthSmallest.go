package _0211029

func kthSmallest(root *TreeNode, k int) int {
	var inorder func(root *TreeNode)
	stack := []int{}
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		stack = append(stack, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return stack[k-1]
}
