package _0210906

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	x := int(321)
	println(Reverse(x))
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	println(MaxArea(height))
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := ThreeSum(nums)
	for i := 0; i < len(res); i++ {
		for j:=0;j<len(res[i]);j++{
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
}
