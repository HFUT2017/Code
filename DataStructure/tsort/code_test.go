package tsort

import "testing"

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	BubbleSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func TestSelectionSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	SelectionSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func TestInsertionSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	InsertionSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func TestHeapSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	HeapSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func TestQuickSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	QuickSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	nums = MergeSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func TestBucketSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	nums = BucketSort(nums, 5)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
