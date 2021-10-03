package _0211003

func maxProfit(prices []int) int {
	if len(prices)==0{
		return 0
	}
	minValue:=prices[0]
	res:=0
	for i:=1;i<len(prices);i++{
		if prices[i]<minValue{
			minValue=prices[i]
		}
		if prices[i]-minValue>res{
			res=prices[i]-minValue
		}
	}
	return res
}
