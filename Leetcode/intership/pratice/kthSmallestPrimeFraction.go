package intership

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	res := [][2]int{}
	for i := 0; i < len(arr)/2; i++ {
		for j := i + 1; j < len(arr); j++ {
			res = append(res, [2]int{arr[i], arr[j]})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return float64(res[i][0])/float64(res[i][1]) < float64(res[j][0])/float64(res[j][1])
	})
	return []int{res[k-1][0], res[k-1][1]}
}
