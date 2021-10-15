package _0211015

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		if sum-root.Val == 0 {
			res++
		}
		dfs(root.Left, sum-root.Val)
		dfs(root.Right, sum-root.Val)
	}
	dfs(root, sum)
	res += pathSum(root.Left, sum)
	res += pathSum(root.Right, sum)
	return res
}
