package main

import (
	"errors"
	"fmt"
)

type segmentTree struct {
	Tree  []int
	Nums  []int
	merge func(v1, v2 int) int
}

func NewSegmentTree(nums []int, merge func(v1, v2 int) int) *segmentTree {
	n := len(nums)
	tree := &segmentTree{
		Tree:  make([]int, n*4),
		Nums:  nums,
		merge: merge,
	}
	tree.buildSegmentTree(0, 0, n-1)
	return tree
}

func (tree *segmentTree) buildSegmentTree(index, l, r int) int {
	if l == r {
		tree.Tree[index] = tree.Nums[l]
		return tree.Nums[l]
	}
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	mid := l + (r-l)/2
	leftTree := tree.buildSegmentTree(leftIndex, l, mid)
	rightTree := tree.buildSegmentTree(rightIndex, mid+1, r)
	tree.Tree[index] = tree.merge(leftTree, rightTree)
	return tree.Tree[index]
}

func (tree *segmentTree) Query(queryL, queryR int) (int, error) {
	n := len(tree.Nums)
	if queryL < 0 || queryR >= n || queryL > queryR {
		return -1, errors.New("please check query interval")
	}
	return tree.queryRange(0, 0, n-1, queryL, queryR), nil
}

func (tree *segmentTree) queryRange(index, l, r, queryL, queryR int) int {
	if l == queryL && r == queryR {
		return tree.Tree[index]
	}
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	mid := l + (r-l)/2
	if queryL > mid {
		return tree.queryRange(rightIndex, mid+1, r, queryL, queryR)
	}
	if queryR <= mid {
		return tree.queryRange(leftIndex, l, mid, queryL, queryR)
	}
	leftTree := tree.queryRange(leftIndex, l, mid, queryL, mid)
	rightTree := tree.queryRange(rightIndex, mid+1, r, mid+1, queryR)
	return tree.merge(leftTree, rightTree)
}

func (tree *segmentTree) Update(key, value int) bool {
	n := len(tree.Nums)
	if key < 0 || key >= n {
		return false
	}
	tree.set(0, 0, n-1, key, value)
	tree.Nums[key] = value
	return true
}

func (tree *segmentTree) set(index, l, r, key, value int) bool {
	if l == r {
		tree.Tree[index] = value
		return true
	}
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2
	mid := l + (r-l)/2
	if key > mid {
		tree.set(rightIndex, mid+1, r, key, value)
	} else {
		tree.set(leftIndex, l, mid, key, value)
	}
	tree.Tree[index] = tree.merge(tree.Tree[leftIndex], tree.Tree[rightIndex])
	return true
}

func multiplication(v1, v2 int) int {
	return v1 + v2
}

func (tree *segmentTree) Print() {
	fmt.Println(tree.Nums)
	fmt.Println(tree.Tree)
}

func main() {
	c := []int{-1, 1, 2, -3, 4, 5, 6}

	a := NewSegmentTree(c, multiplication)
	a.Print()

	resp, err := a.Query(2, 5) //-120
	fmt.Printf("查询结果:%d,  错误:%v\n", resp, err)

	a.Update(2, -2)
	a.Print()
}
