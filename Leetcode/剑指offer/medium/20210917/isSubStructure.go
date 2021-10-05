package _0210917

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSubStructure 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
func IsSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	var res bool
	if A.Val == B.Val {
		res = helper(A, B)
	}
	if !res {
		res = IsSubStructure(A.Left, B)
	}
	if !res {
		res = IsSubStructure(A.Right, B)
	}
	return res
}

func helper(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return helper(A.Left, B.Left) && helper(A.Right, B.Right)
}
