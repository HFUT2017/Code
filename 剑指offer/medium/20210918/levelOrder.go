package _0210918

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
func LevelOrder(root *TreeNode) []int {
	queue := []*TreeNode{}
	index := 0
	queue = append(queue, root)
	size := 1
	res := []int{}
	if root == nil {
		return res
	}
	for index < size {
		n := size - index
		for i := 0; i < n; i++ {
			node := queue[index+i]
			res = append(res, node.Val)
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
	}
	return res
}
