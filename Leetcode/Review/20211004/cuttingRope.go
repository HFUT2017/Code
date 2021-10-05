package _0211004

func cuttingRope(n int) int {
	sum:=1
	if n<4{
		return n-1
	}
	for n>4{
		sum*=3
		n-=3
	}
	return sum*n
}
