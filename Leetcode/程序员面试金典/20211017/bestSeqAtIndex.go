package _0211017

import "sort"

type Person struct {
	height int
	weight int
}

func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	persons := make([]Person, n)
	for i := range persons {
		persons[i].height = height[i]
		persons[i].weight = weight[i]
	}
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].height == persons[j].height {
			return persons[i].weight < persons[j].weight
		}
		return persons[i].height > persons[j].height
	})
	var result []Person
	for _, p := range persons {
		j := sort.Search(len(result), func(i int) bool {
			c := result[i]
			return c.height <= p.height || c.weight <= p.weight
		})
		if j == len(result) {
			result = append(result, p)
		} else {
			result[j] = p
		}
	}
	return len(result)
}
