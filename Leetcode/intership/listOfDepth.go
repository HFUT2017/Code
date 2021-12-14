package intership

func listOfDepth(tree *TreeNode) []*ListNode {
	queue := []*TreeNode{}
	lists := []*ListNode{}
	start, size := 0, 1
	queue = append(queue, tree)
	for start < size {
		n := size - start
		dummyHead := &ListNode{}
		cur := dummyHead
		for i := 0; i < n; i++ {
			node := queue[start+i]
			if node.Left != nil {
				queue = append(queue, node.Left)
				size++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				size++
			}
			cur.Next = &ListNode{Val: node.Val}
			cur = cur.Next
		}
		start += n
		lists = append(lists, dummyHead.Next)
	}
	return lists
}
