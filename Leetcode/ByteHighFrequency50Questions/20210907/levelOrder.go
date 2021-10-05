package _0210907

// LevelOrder 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。
func LevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			res[i] = append(res[i], queue[j].Val)
			if queue[j].Left != nil {
				p = append(p, queue[j].Left)
			}
			if queue[j].Right != nil {
				p = append(p, queue[j].Right)
			}
		}
		queue = p
	}
	return res
}
