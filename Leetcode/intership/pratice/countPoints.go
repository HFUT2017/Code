package intership

func countPoints(rings string) int {
	R := map[int]bool{}
	G := map[int]bool{}
	B := map[int]bool{}
	count := 0
	for i := 0; i < len(rings)-1; i++ {
		if rings[i] == 'R' {
			R[int(rings[i+1]-'0')] = true
		} else if rings[i] == 'G' {
			G[int(rings[i+1]-'0')] = true
		} else {
			B[int(rings[i+1]-'0')] = true
		}
	}
	for index, _ := range R {
		if G[index] == true && B[index] == true {
			count++
		}
	}
	return count
}
