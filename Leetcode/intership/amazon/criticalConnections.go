package amazon

func criticalConnections(n int, connections [][]int) [][]int {
	const maxInt = 1 << 31
	ans := [][]int{}
	parents := make([]int, n)
	time := make([]int, n)
	curTime := 0
	for i, _ := range parents {
		parents[i] = i
		time[i] = maxInt
	}
	graph := make([][]int, n)
	for _, connection := range connections {
		n1, n2 := connection[0], connection[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}
	visited := make([]bool, n)
	var dfsCheckRing func(fromNode, curNode int) int
	dfsCheckRing = func(fromNode, curNode int) int {
		if visited[curNode] && time[curNode] == time[fromNode] {
			return curNode
		}
		if visited[curNode] {
			return -1
		}
		visited[curNode] = true
		time[curNode] = time[fromNode]
		flag := -1
		for _, nextNode := range graph[curNode] {
			if nextNode == fromNode || parents[nextNode] == curNode {
				continue
			}
			if dfsRes := dfsCheckRing(curNode, nextNode); dfsRes != -1 {
				flag = dfsRes
				union(parents, curNode, dfsRes)
			}
		}
		if curNode == flag {
			return -1
		}
		return flag
	}

	for i := 0; i < n; i++ {
		if find(parents, i) == i {
			time[i] = curTime
			dfsCheckRing(i, i)
			curTime++
		}
	}
	for _, connection := range connections {
		n1, n2 := connection[0], connection[1]
		if find(parents, n1) == find(parents, n2) {
			continue
		}
		ans = append(ans, connection)
	}
	return ans
}

func union(parents []int, from, to int) {
	parents[find(parents, from)] = find(parents, to)
}

func find(parents []int, index int) int {
	for parents[index] != index {
		parents[index] = parents[parents[index]]
		index = parents[index]
	}
	return index
}
