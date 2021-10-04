package _0211004

func hammingWeight(num uint32) int {
	var cnt uint32
	cnt=1
	res:=0
	for cnt!=0{
		if cnt&num!=0{
			res++
		}
		cnt<<=1
	}
	return res
}
