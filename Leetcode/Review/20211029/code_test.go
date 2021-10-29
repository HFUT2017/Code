package _0211029

import "testing"

func buildTree_(nums []int, n int) *TreeNode {
	if n >= len(nums) {
		return nil
	}
	root := &TreeNode{Val: nums[n]}
	root.Left = buildTree_(nums, 2*n+1)
	root.Right = buildTree_(nums, 2*n+2)
	return root
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	println(root.Val)
	inorder(root.Right)
}

func TestSumNumbers(t *testing.T) {
	nums := []int{1, 2, 3}
	root := buildTree_(nums, 0)
	inorder(root)
}
