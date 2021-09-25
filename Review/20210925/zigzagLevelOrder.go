package _0210925

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	leftStack := []*TreeNode{}
	rightStack := []*TreeNode{}
	leftStack = append(leftStack, root)
	for len(leftStack) != 0 || len(rightStack) != 0 {
		temp := []int{}

		for len(leftStack) != 0 {
			node := leftStack[len(leftStack)-1]
			temp = append(temp, node.Val)
			if node.Left != nil {
				rightStack = append(rightStack, node.Left)
			}
			if node.Right != nil {
				rightStack = append(rightStack, node.Right)
			}
			leftStack = leftStack[:len(leftStack)-1]
		}

		if len(temp) != 0 {
			res = append(res, temp)
			temp = []int{}
		}

		for len(rightStack) != 0 {
			node := rightStack[len(rightStack)-1]
			temp = append(temp, node.Val)
			if node.Right != nil {
				leftStack = append(leftStack, node.Right)
			}
			if node.Left != nil {
				leftStack = append(leftStack, node.Left)
			}
			rightStack = rightStack[:len(rightStack)-1]
		}

		if len(temp) != 0 {
			res = append(res, temp)
			temp = []int{}
		}
	}
	return res
}
