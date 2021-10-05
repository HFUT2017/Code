package _0210927

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var pathSum func(root *TreeNode) int
	pathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(pathSum(root.Left), 0)
		right := max(pathSum(root.Right), 0)
		res = max(res, left+right+root.Val)
		return max(left, right) + root.Val
	}
	pathSum(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
