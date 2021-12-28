package intership

import "container/heap"

func eatenApples(apples []int, days []int) int {
	h := PriorityQueue{}
	count := 0
	d := 1
	for h.Len() > 0 || d <= len(days) {
		if d <= len(days) {
			if apples[d-1] != 0 && days[d-1] != 0 {
				heap.Push(&h, pair{apples: apples[d-1], days: d - 1 + days[d-1]})
			}
		}
		for h.Len() > 0 && (h[0].apples == 0 || h[0].days < d) {
			heap.Pop(&h)
		}
		if h.Len() != 0 {
			h[0].apples--
			count++
		}
		d++
	}
	return count
}

type pair struct {
	apples int
	days   int
}

type PriorityQueue []pair

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (h PriorityQueue) Less(i, j int) bool { return h[i].days < h[j].days }

func (h *PriorityQueue) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}
