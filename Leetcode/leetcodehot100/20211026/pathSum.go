package _0211026

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rootSum(root *TreeNode, targetSum int) int {
	if root==nil{
		return 0
	}
	res:=0
	if root.Val==targetSum{
		res++
	}
	res+=rootSum(root.Left,targetSum-root.Val)
	res+=rootSum(root.Right,targetSum-root.Val)
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	if root==nil{
		return 0
	}
	res:=rootSum(root,targetSum)
	res+=pathSum(root.Left,targetSum)
	res+=pathSum(root.Right,targetSum)
	return res
}
