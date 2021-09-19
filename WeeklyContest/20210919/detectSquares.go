package _0210919

type DetectSquares struct{}

var row [1001][]int
var col = [1001]map[int]int{}

func Constructor() (_ DetectSquares) {
	row = [1001][]int{}
	for i := range col {
		col[i] = map[int]int{}
	}
	return
}

func (DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	row[y] = append(row[y], x)
	col[x][y]++
}

func (DetectSquares) Count(point []int) (ans int) {
	x, y := point[0], point[1]
	for _, x2 := range row[y] {
		if x2 != x {
			d := abs(x2 - x)
			ans += col[x][y-d] * col[x2][y-d]
			ans += col[x][y+d] * col[x2][y+d]
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
