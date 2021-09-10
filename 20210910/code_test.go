package _0210910

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := Merge(nums)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
}

func TestSearch(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	println(Search(nums, 0))
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := SpiralOrder(matrix)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
}
