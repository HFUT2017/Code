package _0211002

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	index, size := 0, 1
	for index < size {
		n := size - index
		temp := []int{}
		for i := 0; i < n; i++ {
			node := queue[index+i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
				size++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				size++
			}
		}
		index += n
		res = append(res, temp)
	}
	return res
}
