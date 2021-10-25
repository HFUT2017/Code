package _0211025

func hammingDistance(x int, y int) int {
	res := 0
	for i := 0; i < 32; i++ {
		a := (x >> i) & 1
		b := (y >> i) & 1
		res += a ^ b
	}
	return res
}
