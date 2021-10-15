package _0211015

import "math"

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, low, up int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= up {
		return false
	}
	return helper(root.Left, low, root.Val) && helper(root.Right, root.Val, up)
}
