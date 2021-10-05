package _0210923

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	index, size := 0, 1
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
		ans = append(ans, res[len(res)-1])
		index += n
	}
	return ans

}
