package _0211009

//给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符
func minDistance_D(word1 string, word2 string) int {
	res := longestCommonSubsequence(word1, word2)
	return len(word1) + len(word2) - 2*res
}
