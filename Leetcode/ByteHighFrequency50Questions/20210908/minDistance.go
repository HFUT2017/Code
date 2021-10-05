package _0210908

// MinDistance 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
func MinDistance(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(word2)+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			insert := dp[i-1][j] + 1
			delete := dp[i][j-1] + 1
			change := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				change++
			}
			dp[i][j] = min(insert, min(delete, change))
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
