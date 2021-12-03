package intership

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	hash := map[int]bool{}
	hash[firstPerson] = true
	hash[0] = true
	res := []int{}
	for i := 0; i < len(meetings); i++ {
		if hash[meetings[i][0]] == true || hash[meetings[i][1]] == true {
			hash[meetings[i][0]] = true
			hash[meetings[i][1]] = true
		}
	}
	for i := 0; i < n; i++ {
		if hash[i] == true {
			res = append(res, i)
		}
	}
	return res
}
