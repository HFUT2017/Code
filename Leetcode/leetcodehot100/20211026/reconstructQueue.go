package _0211026

import "sort"

func reconstructQueue(people [][]int) [][]int {
	res:=[][]int{}
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	for _, person := range people {
		index := person[1]
		res=append(res[:index],append([][]int{person},res[index:]...)...)
	}
	return res
}
