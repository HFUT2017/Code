package _0211001

func findKthNumber(n int, k int) int {
	var count func(prefix int) int
	count= func(prefix int) int {
		now,next,res:=prefix,prefix+1,0
		for now<=n{
			res+=min(next,n+1)-now
			now*=10
			next*=10
		}
		return res
	}
	prenum:=1
	now:=1
	for prenum<k{
		num:=count(now)
		if prenum+num>k{
			prenum++
			now*=10
		}else{
			prenum+=num
			now++
		}
	}
	return now
}
