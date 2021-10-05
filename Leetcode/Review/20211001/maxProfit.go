package _0211001

import "math"

func maxProfit(prices []int) int {
	min := prices[0]
	res := math.MinInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > res {
			res = prices[i] - min
		}
	}
	return res
}
