package _0211017

func dfs(visit []bool, graph [][]int, start int, target int) bool {
	for i := 0; i < len(graph); i++ {
		if visit[i] {
			continue
		}
		if graph[i][0] == start && graph[i][1] == target {
			return true
		}
		visit[i] = true
		if graph[i][1] == target && dfs(visit, graph, start, graph[i][0]) {
			return true
		}
		visit[i] = false
	}
	return false
}
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	visit := make([]bool, len(graph))
	return dfs(visit, graph, start, target)
}

//邻接表+BFS
//func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
//	hash := make([]map[int]bool, n)
//	for i := 0; i < n; i++ {
//		hash[i] = make(map[int]bool)
//	}
//	for _, edge := range graph {
//		if edge[0] != edge[1] {
//			hash[edge[0]][edge[1]] = true
//		}
//	}
//	queue := []int{start}
//	visit := make([]bool, n)
//	for len(queue) > 0 {
//		cur := queue[0]
//		queue = queue[1:]
//		if visit[cur] == true {
//			continue
//		}
//		if cur == target {
//			return true
//		}
//		for k, _ := range hash[cur] {
//			if visit[k] == false {
//				queue = append(queue, k)
//			}
//		}
//	}
//	return false
//}
