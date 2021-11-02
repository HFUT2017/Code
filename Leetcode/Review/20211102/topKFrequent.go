package _0211102

import "sort"

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}
	res := []int{}
	for _, value := range nums {
		hash[value]++
	}
	for index, _ := range hash {
		res = append(res, index)
	}
	sort.Slice(res, func(i, j int) bool {
		return hash[res[i]] > hash[res[j]]
	})
	return res[:k]
}
