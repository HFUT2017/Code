package _0211012

type LRUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*DLinkedNode
	Head     *DLinkedNode
	Tail     *DLinkedNode
}

type DLinkedNode struct {
	Key   int
	Value int
	Next  *DLinkedNode
	Prev  *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Size:     0,
		Capacity: capacity,
		Cache:    map[int]*DLinkedNode{},
		Head:     &DLinkedNode{Key: 0, Value: 0, Next: nil, Prev: nil},
		Tail:     &DLinkedNode{Key: 0, Value: 0, Next: nil, Prev: nil},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

func (this *LRUCache) DeleteNode(node *DLinkedNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *DLinkedNode) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next = node
	node.Next.Prev = node
}

func (this *LRUCache) MoveToHead(node *DLinkedNode) {
	this.DeleteNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) DeleteTail() *DLinkedNode {
	node := this.Tail.Prev
	node.Prev.Next = this.Tail
	this.Tail.Prev = node.Prev
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

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cache[key]; ok {
		node := this.Cache[key]
		node.Value = value
		this.MoveToHead(node)
	} else {
		node := &DLinkedNode{Key: key, Value: value, Next: nil, Prev: nil}
		this.Cache[key] = node
		this.AddToHead(node)
		this.Size++
		if this.Size > this.Capacity {
			deletedTail := this.DeleteTail()
			delete(this.Cache, deletedTail.Key)
			this.Size--
		}
	}
}
