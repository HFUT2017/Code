package intership

import "sort"

func findEvenNumbers(digits []int) []int {
	hash := map[int]bool{}
	res := []int{}
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(digits); k++ {
				if i == k || j == k {
					continue
				}
				hash[digits[i]*100+digits[j]*10+digits[k]] = true
			}
		}
	}
	for index, _ := range hash {
		if index > 99 && index%2 == 0 {
			res = append(res, index)
		}
	}
	sort.Ints(res)
	return res
}
