package _0210930

type LRUCache struct {
	Capacity int
	Size     int
	Cache    map[int]*DListNode
	Head     *DListNode
	Tail     *DListNode
}

type DListNode struct {
	Val  int
	Key  int
	Next *DListNode
	Prev *DListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Size:     0,
		Cache:    map[int]*DListNode{},
		Head:     &DListNode{Key: 0, Val: 0},
		Tail:     &DListNode{Key: 0, Val: 0},
	}
	lru.Head.Next = lru.Tail
	lru.Head.Prev = lru.Tail
	lru.Tail.Next = lru.Head
	lru.Tail.Prev = lru.Head
	return lru
}

func (this *LRUCache) deleteNode(node *DListNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	node.Next = nil
	node.Prev = nil
}

func (this *LRUCache) deleteTail() *DListNode {
	node := this.Tail.Prev
	node.Prev.Next = this.Tail
	this.Tail.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
	return node
}

func (this *LRUCache) AddToHead(node *DListNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next = node
	node.Next.Prev = node
}

func (this *LRUCache) MoveToHead(node *DListNode) {
	this.deleteNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.MoveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Val = value
		this.MoveToHead(node)
	} else {
		node := &DListNode{Key: key, Val: value}
		this.Cache[key] = node
		this.AddToHead(node)
		this.Size++
		if this.Size > this.Capacity {
			t := this.deleteTail()
			delete(this.Cache, t.Key)
			this.Size--
		}
	}
}
