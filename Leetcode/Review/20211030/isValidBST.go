package _0211030

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var isValid func(root *TreeNode, max int, min int) bool
	isValid = func(root *TreeNode, max int, min int) bool {
		if root == nil {
			return true
		}
		if root.Val >= max || root.Val <= min {
			return false
		}
		return isValid(root.Left, root.Val, min) && isValid(root.Right, max, root.Val)
	}
	return isValid(root, math.MaxInt64, math.MinInt64)
}
