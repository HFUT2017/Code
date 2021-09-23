package _0210922

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	nums := []int{1}
	res := [][]int{}
	res = append(res, nums)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
	nums=append(nums,2)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
	nums=append(nums,2)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
	nums=append(nums,2)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
	nums=append(nums,2)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
}
