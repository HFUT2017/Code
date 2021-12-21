package intership

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	indexA, indexB := 0, len(plants)-1
	nowA, nowB := capacityA, capacityB
	res := 0
	for indexA <= indexB {
		if indexA == indexB {
			if max(nowA, nowB) < plants[indexA] {
				res++
			}
			return res
		}
		if nowA < plants[indexA] && capacityA >= plants[indexA] {
			nowA = capacityA - plants[indexA]
			res++
			indexA++
		} else if nowA >= plants[indexA] {
			nowA -= plants[indexA]
			indexA++
		}
		if nowB < plants[indexB] && capacityB >= plants[indexB] {
			nowB = capacityB - plants[indexB]
			res++
			indexB++
		} else if nowB >= plants[indexB] {
			nowB -= plants[indexB]
			indexB++
		}
		if capacityA < plants[indexA] && capacityB < plants[indexB] {
			return -1
		}
	}
	return 0
}
