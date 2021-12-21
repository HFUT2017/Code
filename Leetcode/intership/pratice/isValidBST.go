package intership

import "math"

func isValidBST(root *TreeNode) bool {
	var judgeBST func(root *TreeNode, maxNUm, minNum int) bool
	judgeBST = func(root *TreeNode, maxNUm, minNum int) bool {
		if root == nil {
			return true
		}
		if root.Val >= maxNUm || root.Val <= minNum {
			return false
		}
		return judgeBST(root.Left, root.Val, minNum) && judgeBST(root.Right, maxNUm, minNum)
	}
	return judgeBST(root, math.MaxInt32, math.MinInt32)
}
