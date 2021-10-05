package _0210922

func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]
	for _, value := range prices {
		if value < min {
			min = value
		}
		if value-min > profit {
			profit = value - min
		}
	}
	return profit
}
