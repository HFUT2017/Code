package intership

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &PriorityQueue{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

//type PriorityQueue [][2]int
//
//func (p PriorityQueue) Len() int { return len(p) }
//
//func (p PriorityQueue) Less(i, j int) bool { return p[i][1] > p[j][1] }
//
//func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
//
//func (p *PriorityQueue) Push(x interface{}) {
//	*p = append(*p, x.([2]int))
//}
//
//func (p *PriorityQueue) Pop() interface{} {
//	n := len(*p)
//	old := *p
//	item := old[n-1]
//	*p = old[:n-1]
//	return item
//}
