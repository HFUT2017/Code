package _0211027

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	hash := map[*Node]*Node{}
	for p := head; p != nil; p = p.Next {
		hash[p] = &Node{Val: p.Val, Next: nil, Random: nil}
	}
	for p := head; p != nil; p = p.Next {
		hash[p].Next = hash[p.Next]
		hash[p].Random = hash[p.Random]
	}
	return hash[head]
}
