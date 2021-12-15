package intership

func LoudAndRich(richer [][]int, quiet []int) []int {
	graph := make([][]int, len(quiet))
	answer := make([]int, len(quiet))
	for i := 0; i < len(richer); i++ {
		graph[richer[i][1]] = append(graph[richer[i][1]], richer[i][0])
	}
	for i := 0; i < len(quiet); i++ {
		answer[i] = -1
	}
	var dfs func(index int)
	dfs = func(index int) {
		if answer[index] != -1 {
			return
		}
		answer[index] = index
		for _, rich := range graph[index] {
			dfs(rich)
			if quiet[answer[rich]] < quiet[answer[index]] {
				answer[index] = answer[rich]
			}
		}
	}
	for i := 0; i < len(quiet); i++ {
		dfs(i)
	}
	return answer
}

func loudAndRich(richer [][]int, quiet []int) []int {
	personNum := len(quiet)
	queue := make([]int, 0)
	answer := make([]int, personNum)
	inDegree := make([]int, personNum)
	graph := make([][]int, personNum)
	for _, rich := range richer {
		graph[rich[0]] = append(graph[rich[0]], rich[1])
		inDegree[rich[1]] += 1
	}
	for index, _ := range quiet {
		answer[index] = index
		if inDegree[index] == 0 {
			queue = append(queue, index)
		}
	}
	for len(queue) > 0 {
		personRicher := queue[0]
		queue = quiet[1:]
		for _, person := range graph[personRicher] {
			if quiet[personRicher] < quiet[person] {
				answer[person] = personRicher
			}
			inDegree[person] -= 1
			if inDegree[person] == 0 {
				queue = append(queue, person)
			}
		}
	}
	return answer
}
