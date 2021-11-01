package _0211031

func minimumOperations(nums []int, start, goal int) int {
	used := [1001]bool{}
	queue := []int{start}
	used[start] = true
	for step := 1; len(queue) > 0; step++ {
		temp := queue
		queue = nil
		for _, value := range temp {
			for _, num := range nums {
				for _, res := range []int{value + num, value - num, value ^ num} {
					if res == goal {
						return step
					}
					if res >= 0 && res <= 1000 && used[res] == false {
						used[res] = true
						queue = append(queue, res)
					}
				}
			}
		}
	}
	return -1
}
