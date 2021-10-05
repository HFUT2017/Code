package _0210915

func count(x int) int {
	sum := 0
	for x != 0 {
		sum = sum + x%10
		x /= 10
	}
	return sum
}

func check(m, n, k int, i, j int) bool {
	if i < 0 || i >= m || j < 0 || j >= n || (count(i)+count(j) > k) {
		return false
	}
	return true
}

func MovingCount(m int, n int, k int) int {
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	x, y := 0, 0
	var backTrack func() int
	backTrack = func() int {
		if check(m, n, k, x, y) == false || used[x][y] == true {
			return 0
		}
		used[x][y] = true
		sum := 1
		for i := 0; i < len(dirs); i++ {
			x = x + dirs[i][0]
			y = y + dirs[i][1]
			sum += backTrack()
		}
		return sum
	}
	return backTrack()
}
