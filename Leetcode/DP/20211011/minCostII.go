package _0211011

import "math"

//func minCostII(costs [][]int) int {
//	dp := make([][]int, len(costs))
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, len(costs[0]))
//	}
//	for i := 0; i < len(costs[0]); i++ {
//		dp[0][i] = costs[0][i]
//	}
//	for i := 1; i < len(costs); i++ {
//		for j := 0; j < len(costs[i]); j++ {
//			dp[i][j] = math.MaxInt32
//			for k := 0; k < len(costs[i]); k++ {
//				if k != j {
//					dp[i][j] = min(dp[i][j], dp[i-1][k]+costs[i][j])
//				}
//			}
//		}
//	}
//	res := math.MaxInt32
//	for i := 0; i < len(costs[0]); i++ {
//		res = min(dp[len(costs)-1][i], res)
//	}
//	return res
//}

func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	k := len(costs[0])
	if k < 2 {
		return costs[0][0]
	}
	min, secondMin, minColor := 0, 0, 0
	for _, cost := range costs {
		tMin, tSecond, tMinColor := math.MaxInt32, math.MaxInt32, -1
		for i := 0; i < k; i++ {
			var minCost int
			if i == minColor {
				minCost = secondMin + cost[i]
			} else {
				minCost = min + cost[i]
			}
			if minCost < tMin {
				tSecond = tMin
				tMin = minCost
				tMinColor = i
			} else if minCost < tSecond {
				tSecond = minCost
			}
		}
		min, secondMin, minColor = tMin, tSecond, tMinColor
	}
	return min
}
