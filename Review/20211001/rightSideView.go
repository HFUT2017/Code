package _0211001

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	ans := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	index, size := 0, 1
	for index < size {
		n := size - index
		for i := 0; i < n; i++ {
			node := queue[index+i]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
				size++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				size++
			}
		}
		res = append(res, ans[len(ans)-1])
		index += n
	}
	return res
}
