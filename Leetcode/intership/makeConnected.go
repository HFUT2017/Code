package intership

func MakeConnected(n int, connections [][]int) int {
	parent := make([]int, n)
	count := 0
	ans := 0
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(from, to int)
	union = func(from, to int) {
		if find(from) == find(to) {
			count++
		}
		parent[find(from)] = find(to)
	}
	for _, value := range connections {
		union(value[0], value[1])
	}
	for i := 0; i < n; i++ {
		if parent[i] == i {
			ans++
		}
	}
	if ans > (count + 1) {
		return -1
	}
	return ans - 1
}
