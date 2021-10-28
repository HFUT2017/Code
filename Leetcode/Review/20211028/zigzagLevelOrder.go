package _0211028

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	left, right := []*TreeNode{root}, []*TreeNode{}
	res := [][]int{}
	for len(left) != 0 || len(right) != 0 {
		temp := []int{}
		if len(left) != 0 {
			for len(left) != 0 {
				node := left[len(left)-1]
				left = left[:len(left)-1]
				temp = append(temp, node.Val)
				if node.Left != nil {
					right = append(right, node.Left)
				}
				if node.Right != nil {
					right = append(right, node.Right)
				}
			}
		} else {
			for len(right) != 0 {
				node := right[len(right)-1]
				right = right[:len(right)-1]
				temp = append(temp, node.Val)
				if node.Right != nil {
					left = append(left, node.Right)
				}
				if node.Left != nil {
					left = append(left, node.Left)
				}
			}
		}
		res = append(res, temp)
	}
	return res
}
