package _0211102

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for index, _ := range dp {
		dp[index] = math.MaxInt32
	}
	dp[0] = 0
	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
				continue
			}
			break
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
