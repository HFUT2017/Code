package amazon

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	left, right := []*TreeNode{root}, []*TreeNode{}
	for len(left) > 0 || len(right) > 0 {
		t := []int{}
		if len(left) > 0 {
			n := len(left)
			for i := 0; i < n; i++ {
				node := left[len(left)-1]
				t = append(t, node.Val)
				left = left[:len(left)-1]
				if node.Left != nil {
					right = append(right, node.Left)
				}
				if node.Right != nil {
					right = append(right, node.Right)
				}
			}
		} else {
			n := len(right)
			for i := 0; i < n; i++ {
				node := right[len(right)-1]
				t = append(t, node.Val)
				right = right[:len(right)-1]
				if node.Right != nil {
					left = append(left, node.Right)
				}
				if node.Left != nil {
					left = append(left, node.Left)
				}
			}
		}
		res = append(res, t)
	}
	return res
}
