package _0211001

func maxPathSum(root *TreeNode) int {
	res := 0
	var maxPath func(node *TreeNode) int
	maxPath = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(maxPath(node.Left), 0)
		right := max(maxPath(node.Right), 0)
		res = max(left+right+1, res)
		return max(left, right) + 1
	}
	maxPath(root)
	return res
}
