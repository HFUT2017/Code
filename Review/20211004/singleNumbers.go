package _0211004

func singleNumbers(nums []int) []int {
	sum:=0
	a,b:=0,0
	for _,value:=range nums{
		sum^=value
	}
	cnt:=1
	for cnt&sum==0{
		cnt<<=1
	}
	for _,value:=range nums{
		if cnt&value==0{
			a^=value
		}else{
			b^=value
		}
	}
	return []int{a,b}
}
