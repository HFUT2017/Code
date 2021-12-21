package amazon

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	p := PriorityQueue{}
	res := [][]int{}
	for _, point := range points {
		heap.Push(&p, pair{x: point[0], y: point[1]})
		if p.Len() > k {
			heap.Pop(&p)
		}
	}
	for p.Len() > 0 {
		item := heap.Pop(&p)
		res = append(res, []int{item.(pair).x, item.(pair).y})
	}
	return res
}

type pair struct{ x, y int }

type PriorityQueue []pair

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].x*p[i].x+p[i].y*p[i].y > p[j].x*p[j].x+p[j].y*p[j].y
}

func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *PriorityQueue) Push(x interface{}) { *p = append(*p, x.(pair)) }

func (p *PriorityQueue) Pop() interface{} {
	n := len(*p)
	old := *p
	item := old[n-1]
	*p = old[:n-1]
	return item
}
