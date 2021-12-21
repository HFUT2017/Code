package amazon

type Node struct {
	Val    *Node
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hash := map[*Node]*Node{}
	cur := head
	for cur != nil {
		hash[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur = cur.Next
	}
	return hash[head]
}
