package amazon

import "sort"

func topKFrequent(words []string, k int) []string {
	hash := map[string]int{}
	for _, word := range words {
		hash[word]++
	}
	if k > len(hash) {
		return []string{}
	}
	str := []string{}
	for key, _ := range hash {
		str = append(str, key)
	}
	sort.Slice(str, func(i, j int) bool {
		if hash[str[i]] == hash[str[j]] {
			return str[i] < str[j]
		}
		return hash[str[i]] > hash[str[j]]
	})
	return str[:k]
}
