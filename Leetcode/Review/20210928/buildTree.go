package _0210928

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := 0
	for index < len(inorder) {
		if inorder[index] == preorder[0] {
			break
		}
		index++
	}
	node := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:len(inorder[:index])+1], inorder[:index]),
		Right: buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:]),
	}
	return node
}
