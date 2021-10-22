package _0211022

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var valid func(root *TreeNode, max, min int) bool
	valid = func(root *TreeNode, max, min int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return valid(root.Left, root.Val, min) && valid(root.Right, max, root.Val)
	}
	return valid(root, math.MaxInt64, math.MinInt64)
}
