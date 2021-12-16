package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue struct {
	Value    string
	Priority int
	Index    int
}

type pqueue []*PriorityQueue

func (pq pqueue) Len() int {
	return len(pq)
}

func (pq pqueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *pqueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueue)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *pqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *pqueue) Update(item *PriorityQueue, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func main() {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(pqueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &PriorityQueue{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)
	item := &PriorityQueue{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 5)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityQueue)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
}
