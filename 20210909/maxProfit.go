package _0210909

import "math"

// MaxProfit 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
func MaxProfit(prices []int) int {
	min, maxProfit := math.MaxInt32, 0
	for _, value := range prices {
		if value < min {
			min = value
		}
		if value-min > maxProfit {
			maxProfit = value - min
		}
	}
	return maxProfit
}
