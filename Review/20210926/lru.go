package _0210926

type LRUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*DLinkNode
	Head     *DLinkNode
	Tail     *DLinkNode
}

type DLinkNode struct {
	Key   int
	Value int
	Prev  *DLinkNode
	Next  *DLinkNode
}

func Constructor(maxSize int) LRUCache {
	lru := LRUCache{
		Size:     maxSize,
		Capacity: 0,
		Cache:    make(map[int]*DLinkNode),
		Head:     &DLinkNode{Key: 0, Value: 0},
		Tail:     &DLinkNode{Key: 0, Value: 0},
	}
	lru.Head.Prev = lru.Tail
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	lru.Tail.Next = lru.Head
	return lru
}

func (this *LRUCache) AddToHead(node *DLinkNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	node.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) DeleteNode(node *DLinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (this *LRUCache) MoveToHead(node *DLinkNode) {
	this.DeleteNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) DeleteTail() *DLinkNode {
	node := this.Tail.Prev
	this.DeleteNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}
	node := this.Cache[key]
	this.MoveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.Cache[key]; ok {
		node := this.Cache[key]
		node.Value = value
		this.MoveToHead(node)
	} else {
		node := &DLinkNode{Key: key, Value: value}
		this.Cache[key] = node
		this.Capacity = this.Capacity + 1
		this.AddToHead(node)
		if this.Capacity > this.Size {
			node := this.DeleteTail()
			delete(this.Cache, node.Key)
			this.Capacity = this.Capacity - 1
		}
	}
}
