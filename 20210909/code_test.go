package _0210909

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{1, 2, 0}
	println(FirstMissingPositive(nums))
}

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(MaxProfit(prices))
}

func TestMaxProfit_2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(MaxProfit_2(prices))
}

func TestNumIslands(t *testing.T) {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	println(NumIslands(grid))
}
