package intership

type NumArray struct {
	Tree   []int
	Nums   []int
	merger func(v1, v2 int) int
}

func (this *NumArray) buildTree(index, l, r int) int {
	if l == r {
		this.Tree[index] = this.Nums[l]
		return this.Nums[l]
	}
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	mid := l + (r-l)/2
	leftTree := this.buildTree(leftIndex, l, mid)
	rightTree := this.buildTree(rightIndex, mid+1, r)
	this.Tree[index] = this.merger(leftTree, rightTree)
	return this.Tree[index]
}
func SegmentTreeConstructor(nums []int) NumArray {
	merger := func(v1, v2 int) int {
		return v1 + v2
	}
	tree := NumArray{
		Tree:   make([]int, len(nums)*2),
		Nums:   nums,
		merger: merger,
	}
	tree.buildTree(0, 0, len(nums)-1)
	return tree
}

func (this *NumArray) set(index, l, r, key, value int) {
	if l == r {
		this.Tree[index] = value
		return
	}
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	mid := l + (r-l)/2
	if key > mid {
		this.set(rightIndex, mid+1, r, key, value)
	} else {
		this.set(leftIndex, l, mid, key, value)
	}
	this.Tree[index] = this.merger(this.Tree[leftIndex], this.Tree[rightIndex])
}

func (this *NumArray) Update(index int, val int) {
	n := len(this.Nums)
	if index < 0 || index > n {
		return
	}
	this.set(0, 0, n-1, index, val)
}

func (this *NumArray) queryRange(index, l, r, left, right int) int {
	if l == left && r == right {
		return this.Tree[index]
	}
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	mid := l + (r-l)/2
	if left > mid {
		return this.queryRange(rightIndex, mid+1, r, left, right)
	}
	if right <= mid {
		return this.queryRange(leftIndex, l, mid, left, right)
	}
	leftTree := this.queryRange(leftIndex, l, mid, left, mid)
	rightTree := this.queryRange(rightIndex, mid+1, r, mid+1, right)
	return this.merger(leftTree, rightTree)
}

func (this *NumArray) SumRange(left int, right int) int {
	n := len(this.Nums)
	if left < 0 || right >= n || left > right {
		return -1
	}
	return this.queryRange(0, 0, n-1, left, right)
}
