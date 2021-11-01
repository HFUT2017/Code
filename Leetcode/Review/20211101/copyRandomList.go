package _0211101

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hash := map[*Node]*Node{}
	cur := head
	for cur != nil {
		t := &Node{Val: cur.Val, Next: nil, Random: nil}
		hash[cur] = t
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur=cur.Next
	}
	return hash[head]
}
