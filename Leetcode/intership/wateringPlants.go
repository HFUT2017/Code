package intership

func wateringPlants(plants []int, capacity int) int {
	res := 0
	now := capacity
	for index, value := range plants {
		if now < value {
			res += index
			now = capacity - value
		} else {
			now -= value
			res += 1
		}
	}
	return res
}
