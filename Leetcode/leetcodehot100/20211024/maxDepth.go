package _0211024

func maxDepth(root *TreeNode) int {
	res := 0
	var getDepth func(root *TreeNode) int
	getDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getDepth(root.Left)
		right := getDepth(root.Right)
		depth := max(left, right) + 1
		res = max(res, depth)
		return depth
	}
	getDepth(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
