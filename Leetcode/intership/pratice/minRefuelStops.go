package intership

import "container/heap"

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	p := &PriorityQueue{}
	count := 0
	index := 0
	heap.Init(p)
	for count = 0; startFuel < target; count++ {
		for index < len(stations) && stations[index][0] <= startFuel {
			heap.Push(p, [2]int{stations[index][0], stations[index][1]})
			index++
		}
		if p.Len() == 0 {
			return -1
		}
		startFuel += heap.Pop(p).([2]int)[1]
	}
	return count
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
