package _0211029

import "strconv"

func SumNumbers(root *TreeNode) int {
	res := 0
	var inorder func(root *TreeNode, temp string)
	inorder = func(root *TreeNode, temp string) {
		if root == nil {
			return
		}
		temp += strconv.Itoa(root.Val)
		inorder(root.Left, temp)
		if root.Left == nil && root.Right == nil {
			num, _ := strconv.Atoi(temp)
			res += num
		}
		inorder(root.Right, temp)
		temp = temp[:len(temp)-1]
	}
	inorder(root, "")
	return res
}
