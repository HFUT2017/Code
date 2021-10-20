package _0211020

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{}
	index, size := 0, 1
	for index < size {
		temp := []int{}
		n := size - index
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
		res = append(res, temp[len(temp)-1])
	}
	return res
}
