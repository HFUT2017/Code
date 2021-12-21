package intership

import (
	"container/heap"
	"sort"
)

type pair struct{ right, height int }
type PriorityQueue []pair

func (p PriorityQueue) Len() int            { return len(p) }
func (p PriorityQueue) Less(i, j int) bool  { return p[i].height > p[j].height }
func (p PriorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Push(x interface{}) { *p = append(*p, x.(pair)) }
func (p *PriorityQueue) Pop() interface{} {
	old := *p
	item := old[len(old)-1]
	*p = old[:len(old)-1]
	return item
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	ans := [][]int{}
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	index := 0
	h := PriorityQueue{}
	for _, boundary := range boundaries {
		for index < n && buildings[index][0] <= boundary {
			heap.Push(&h, pair{buildings[index][1], buildings[index][2]})
			index++
		}
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}
		maxHeight := 0
		if len(h) > 0 {
			maxHeight = h[0].height
		}
		if len(ans) == 0 || maxHeight != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxHeight})
		}
	}
	return ans
}
