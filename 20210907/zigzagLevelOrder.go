package _0210907

// ZigzagLevelOrder 给定一个二叉树，返回其节点值的锯齿形层序遍历。
func ZigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var node *TreeNode
	var temp []int
	var stackLeft, stackRight []*TreeNode
	stackLeft = append(stackLeft, root)
	for len(stackLeft) > 0 || len(stackRight) > 0 {
		for len(stackLeft) > 0 {
			node = stackLeft[len(stackLeft)-1]
			temp = append(temp, node.Val)
			if node.Left != nil {
				stackRight = append(stackRight, node.Left)
			}
			if node.Right != nil {
				stackRight = append(stackRight, node.Right)
			}
			stackLeft = stackLeft[:len(stackLeft)-1]
		}
		if len(temp) != 0 {
			res = append(res, temp)
			temp = []int{}
		}
		for len(stackRight) > 0 {
			node = stackRight[len(stackRight)-1]
			temp = append(temp, node.Val)
			if node.Right != nil {
				stackLeft = append(stackLeft, node.Right)
			}
			if node.Left != nil {
				stackLeft = append(stackLeft, node.Left)
			}
			stackRight = stackRight[:len(stackRight)-1]
		}
		if len(temp) != 0 {
			res = append(res, temp)
			temp = []int{}
		}
	}
	return res
}
