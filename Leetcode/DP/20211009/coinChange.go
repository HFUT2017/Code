package _0211009

//func coinChange(coins []int, amount int) int {
//	dp := make([]int, amount+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = math.MaxInt32
//	}
//	dp[0] = 0
//	for i := 0; i <= amount; i++ {
//		for j := 0; j < len(coins); j++ {
//			if i >= coins[j] {
//				dp[i] = min(dp[i], dp[i-coins[j]]+1)
//			}
//		}
//	}
//	if dp[amount] == math.MaxInt32 {
//		return -1
//	}
//	return dp[amount]
//}

//func CoinChange(coins []int, amount int) int {
//	return dp(coins, amount)
//}
//
//func dp(coins []int, n int) int {
//	if n == 0 {
//		return 0
//	} else if n < 0 {
//		return -1
//	}
//	res := math.MaxInt32
//	for _, value := range coins {
//		sub := dp(coins, n-value)
//		if sub == -1 {
//			continue
//		}
//		res = min(res, sub+1)
//	}
//	if res == math.MaxInt32 {
//		return -1
//	}
//	return res
//}

//func CoinChange(coins []int, amount int) int {
//	mem := make([]int, amount+1)
//	return dp(coins, amount, mem)
//}
//
//func dp(coins []int, n int, mem []int) int {
//	if n == 0 {
//		return 0
//	} else if n < 0 {
//		return -1
//	} else if mem[n] != 0 {
//		return mem[n]
//	}
//	res := math.MaxInt32
//	for _, value := range coins {
//		sub := dp(coins, n-value, mem)
//		if sub == -1 {
//			continue
//		}
//		res = min(res, sub+1)
//	}
//	if res == math.MaxInt32 {
//		return -1
//	}
//	mem[n] = res
//	return mem[n]
//}
