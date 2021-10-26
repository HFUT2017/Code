package _0211027

type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		nums: []int{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.nums) == 0 {
		this.nums = append(this.nums, num)
	} else {
		i := 0
		for this.nums[i] < num {
			i++
		}
		temp := this.nums[:i]
		temp = append(temp, num)
		temp = append(temp, this.nums[i:]...)
		this.nums = temp
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l := len(this.nums)
	if l%2 == 1 {
		return float64(this.nums[l/2])
	} else {
		return float64(this.nums[l/2-1]+this.nums[l/2]) / 2.0
	}
}
