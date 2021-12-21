package intership

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	ans := 0
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(from, to int)
	union = func(from, to int) {
		parent[find(from)] = find(to)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if parent[i] == i {
			ans++
		}
	}
	return ans
}
