package everyday

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	hash := map[int]bool{}
	var inorder func(root *TreeNode)
	res := false
	inorder = func(root *TreeNode) {
		if root == nil || res == true {
			return
		}
		if hash[k-root.Val] == true {
			res = true
		}
		hash[root.Val] = true
		inorder(root.Left)
		inorder(root.Right)
	}
	inorder(root)
	return res
}
