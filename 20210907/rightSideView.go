package _0210907

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// RightSideView 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
func RightSideView(root *TreeNode) []int {
	res := [][]int{}
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		queue = p
	}
	ret := []int{}
	for i := 0; i < len(res); i++ {
		ret = append(ret, res[i][len(res[i])-1])
	}
	return ret
}
