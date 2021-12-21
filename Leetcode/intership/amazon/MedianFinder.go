package amazon

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }

func (h *hp) Pop() interface{} {
	old := h.IntSlice
	v := old[len(old)-1]
	h.IntSlice = old[:len(old)-1]
	return v
}

type MedianFinder struct {
	qMin, qMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	qMin, qMax := &this.qMin, &this.qMax
	if qMin.Len() == 0 || num <= -qMin.IntSlice[0] {
		heap.Push(qMin, -num)
		if qMax.Len()+1 < qMin.Len() {
			heap.Push(qMax, -heap.Pop(qMin).(int))
		}
	} else {
		heap.Push(qMax, num)
		if qMax.Len() > qMin.Len() {
			heap.Push(qMin, -heap.Pop(qMax).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	qMin, qMax := &this.qMin, &this.qMax
	if qMin.Len() > qMax.Len() {
		return float64(-qMin.IntSlice[0])
	} else {
		return float64(qMax.IntSlice[0]-qMin.IntSlice[0]) / 2.0
	}
}
