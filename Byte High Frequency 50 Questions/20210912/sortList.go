package _0210912

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	mid := findMiddleSortList(head)
	l1 := SortList(head)
	l2 := SortList(mid)
	return merge(l1, l2)
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			head.Next = head1
			head = head.Next
			head1 = head1.Next
		} else {
			head.Next = head2
			head = head.Next
			head2 = head2.Next
		}
	}
	if head1 == nil {
		head.Next = head2
	}
	if head2 == nil {
		head.Next = head1
	}
	return node.Next
}

func findMiddleSortList(node *ListNode) *ListNode {
	slow := node
	fast := node.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	t := slow.Next
	slow.Next = nil
	return t
}
