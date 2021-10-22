package _0211023

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := helper(start, i-1)
		rightTrees := helper(i+1, end)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = leftTree
				root.Right = rightTree
				res = append(res, root)
			}
		}
	}
	return res
}
