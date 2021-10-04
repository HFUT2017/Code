package _0211004

func cuttingRope2(n int) int {
	sum:=1
	if n<4{
		return n-1
	}
	for n>4{
		sum=(sum*3)%1000000007
		n-=3
	}
	return (sum*n)%1000000007
}
