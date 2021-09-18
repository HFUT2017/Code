package _0210918

import "math"

// MaxProfit 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
func MaxProfit(prices []int) int {
	min, res := math.MaxInt32, 0
	for _, value := range prices {
		if value < min {
			min = value
		}
		if value-min > res {
			res = value - min
		}
	}
	return res
}
