package _0211031

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		selected := root.Val + left[1] + right[1]
		unselected := max(left[0], left[1]) + max(right[0], right[1])
		return []int{selected, unselected}
	}
	res := dfs(root)
	return max(res[0], res[1])
}
