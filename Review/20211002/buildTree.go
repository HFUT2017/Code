package _0211002

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	index := 0
	for inorder[index] != preorder[0] {
		index++
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
