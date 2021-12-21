package amazon

func partitionLabels(s string) []int {
	hash := map[byte]int{}
	res := []int{}
	for index, _ := range s {
		hash[s[index]] = index
	}
	start := 0
	for start < len(s) {
		end := hash[s[start]]
		for j := start; j < end; j++ {
			end = max(end, hash[s[j]])
		}
		res = append(res, end-start+1)
		start = end + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
