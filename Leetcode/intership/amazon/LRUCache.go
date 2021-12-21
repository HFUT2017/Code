package amazon

type DListNode struct {
	Key  int
	Val  int
	Next *DListNode
	Prev *DListNode
}
type LRUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*DListNode
	Head     *DListNode
	Tail     *DListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Size:     0,
		Capacity: capacity,
		Cache:    map[int]*DListNode{},
		Head:     &DListNode{Key: 0, Val: 0},
		Tail:     &DListNode{Key: 0, Val: 0},
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	return lru
}

func (this *LRUCache) deleteNode(node *DListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *DListNode) {
	node.Next = this.Head.Next
	node.Next.Prev = node
	this.Head.Next = node
	node.Prev = this.Head
}

func (this *LRUCache) MoveToHead(node *DListNode) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.MoveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) deleteTail() *DListNode {
	node := this.Tail.Prev
	this.deleteNode(node)
	return node
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Val = value
		this.MoveToHead(node)
	} else {
		node := &DListNode{Key: key, Val: value}
		this.Cache[key] = node
		this.addToHead(node)
		this.Size++
		if this.Size > this.Capacity {
			t := this.deleteTail()
			delete(this.Cache, t.Key)
			this.Size--
		}
	}
}
