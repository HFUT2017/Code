package intership

import (
	"math"
	"sort"
)

func VisiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	angles := []float64{}
	for _, point := range points {
		if point[0] == location[0] && point[1] == location[1] {
			sameCnt++
		} else {
			angles = append(angles, math.Atan2(float64(point[1]-location[1]), float64(point[0]-location[0])))
		}
	}
	sort.Float64s(angles)
	n := len(angles)
	for _, degree := range angles {
		angles = append(angles, degree+2*math.Pi)
	}
	count := 0
	right := 0
	degree := float64(angle) * math.Pi / 180
	for index, value := range angles[:n] {
		for right < n*2 && angles[right] < value+degree {
			right++
		}
		if right-index > count {
			count = right - index
		}
	}
	return count + sameCnt
}
