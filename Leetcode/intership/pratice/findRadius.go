package intership

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	j := 0
	res := 0
	for i := 0; i < len(houses); i++ {
		dis := abs(heaters[j] - houses[i])
		for (j+1) < len(heaters) && abs(heaters[j]-houses[i]) >= abs(heaters[j+1]-houses[i]) {
			j++
			if abs(heaters[j]-houses[i]) < dis {
				dis = abs(heaters[j] - houses[i])
			}
		}
		if dis > res {
			res = dis
		}
	}
	return res
}
