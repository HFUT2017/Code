package amazon

func orangesRotting(grid [][]int) int {
	queue := [][]int{}
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	start, size, count := 0, len(queue), 0
	for start < size {
		n := size - start
		for i := 0; i < n; i++ {
			point := queue[i+start]
			for _, d := range dir {
				x, y := point[0]+d[0], point[1]+d[1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				queue = append(queue, []int{x, y})
				size++
			}
		}
		count++
		start += n
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if size == 0 {
		return 0
	}
	return count - 1
}
