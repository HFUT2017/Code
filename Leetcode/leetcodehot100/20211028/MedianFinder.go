package _0211028

import "sort"

type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		nums: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.nums)
	l := len(this.nums)
	if l%2 == 1 {
		return float64(this.nums[l/2])
	} else {
		return float64(this.nums[l/2-1]+this.nums[l/2]) / 2.0
	}
}
