package tsearch

import (
	"sort"
	"testing"
)

func TestOrderSearch(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	println(OrderSearch(nums, 3))
}

func TestBinarySearch(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	println(BinarySearch(nums, 3))
}

func TestBinarySearch_(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	println(BinarySearch_(nums, 3, 0, len(nums)))
}

func TestInsertSearch(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	sort.Ints(nums)
	println(InsertSearch(nums,5))
}
