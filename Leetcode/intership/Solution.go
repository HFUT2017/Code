package intership

import "math/rand"

type Solution struct {
	row    int
	col    int
	matrix map[[2]int]bool
}

func Constructor(m int, n int) Solution {
	return Solution{
		row:    m,
		col:    n,
		matrix: map[[2]int]bool{},
	}

}

func (this *Solution) Flip() []int {
	for {
		row, col := rand.Intn(this.row), rand.Intn(this.col)
		_, ok := this.matrix[[2]int{row, col}]
		if !ok {
			return []int{row, col}
		}
	}
}

func (this *Solution) Reset() {
	this.matrix = map[[2]int]bool{}
}
