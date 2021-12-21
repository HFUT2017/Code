package amazon

import (
	"container/heap"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	var p hp
	count := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		if p.Len() == 0 {
			heap.Push(&p, interval[1])
			count++
		} else {
			if p.IntSlice[0] > interval[0] {
				count++
			} else {
				heap.Pop(&p)
			}
			heap.Push(&p, interval[1])
		}
	}
	return count
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	old := h.IntSlice
	v := old[len(old)-1]
	h.IntSlice = old[:len(old)-1]
	return v
}
