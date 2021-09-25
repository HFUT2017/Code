package _0210911

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(maxGain(node.Left), 0)
		right := max(maxGain(node.Right), 0)
		sum := left + right + node.Val
		maxSum = max(sum, maxSum)
		return node.Val + max(left, right)
	}
	maxGain(root)
	return maxSum
}
