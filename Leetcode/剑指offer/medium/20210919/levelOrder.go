package _0210919

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	level := 0
	for len(queue) != 0 {
		temp := []*TreeNode{}
		res = append(res, []int{})
		if level%2 == 0 {
			for _, value := range queue {
				res[level] = append(res[level], value.Val)
			}
			for i := len(queue) - 1; i >= 0; i-- {
				node := queue[i]
				if node.Right != nil {
					temp = append(temp, node.Right)
				}
				if node.Left != nil {
					temp = append(temp, node.Left)
				}
			}
		} else {
			for _, value := range queue {
				res[level] = append(res[level], value.Val)
			}
			for i := len(queue) - 1; i >= 0; i-- {
				node := queue[i]
				if node.Left != nil {
					temp = append(temp, node.Left)
				}
				if node.Right != nil {
					temp = append(temp, node.Right)
				}
			}
		}
		level++
		queue = temp
	}
	return res
}
