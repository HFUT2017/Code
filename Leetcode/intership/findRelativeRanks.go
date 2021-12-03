package intership

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	hash := map[int]int{}
	res := []string{}
	score_temp := make([]int, len(score))
	copy(score_temp, score)
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})
	for index, value := range score {
		hash[value] = index + 1
	}
	for _, value := range score_temp {
		if hash[value] == 1 {
			res = append(res, "Gold Medal")
		} else if hash[value] == 2 {
			res = append(res, "Silver Medal")
		} else if hash[value] == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(hash[value]))
		}
	}
	return res
}
