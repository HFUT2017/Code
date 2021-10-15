package _0211015

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	hash := map[string][]int{}
	for index, value := range words {
		hash[value] = append(hash[value], index)
	}
	res := math.MaxInt32
	for i, j := 0, 0; i < len(hash[word1]) && j < len(hash[word2]); {
		if hash[word1][i] > hash[word2][j] {
			res = min(res, hash[word1][i]-hash[word2][j])
			j++
		} else {
			res = min(res, hash[word2][j]-hash[word1][i])
			i++
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
