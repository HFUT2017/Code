package _0211020

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)
		res = max(res, left+right+root.Val)
		return max(left, right) + root.Val
	}
	dfs(root)
	return res
}
