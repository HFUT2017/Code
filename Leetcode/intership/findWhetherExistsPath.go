package intership

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	dict := map[int][]int{}
	queue := []int{}
	for i := 0; i < n; i++ {
		dict[i] = []int{}
	}
	for i := 0; i < len(graph); i++ {
		dict[graph[i][0]] = append(dict[graph[i][0]], graph[i][1])
	}
	queue = append(queue, start)
	for len(queue) != 0 {
		now := queue[0]
		queue = queue[1:]
		if now == target {
			return true
		}
		for _, value := range dict[now] {
			queue = append(queue, value)
		}
	}
	return false
}
