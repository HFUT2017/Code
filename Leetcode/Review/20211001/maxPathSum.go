package _0211001

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var maxPath func(node *TreeNode) int
	maxPath = func(node *TreeNode) int {
		if node==nil{
			return 0
		}
		left := max(maxPath(node.Left), 0)
		right := max(maxPath(node.Right), 0)
		res = max(left+right+node.Val, res)
		return max(left, right) + node.Val
	}
	maxPath(root)
	return res
}
