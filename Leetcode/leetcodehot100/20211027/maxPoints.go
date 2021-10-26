package _0211027

func maxPoints(points [][]int) int {
	res := 1
	for i := 0; i < len(points); i++ {
		x := points[i]
		cnt := 1
		for j := i + 1; j < len(points); j++ {
			y := points[j]
			cnt = 2
			for k := j + 1; k < len(points); k++ {
				z := points[k]
				if (x[0]-y[0])*(y[1]-z[1]) == (y[0]-z[0])*(x[1]-y[1]) {
					cnt++
				}
			}
			res = max(res, cnt)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
