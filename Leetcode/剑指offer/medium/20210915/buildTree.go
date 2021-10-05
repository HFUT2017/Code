package _0210915

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:len(inorder[:index])+1], inorder[:index]),
		Right: BuildTree(preorder[len(inorder[:index])+1:], inorder[index+1:]),
	}
	return root
}
