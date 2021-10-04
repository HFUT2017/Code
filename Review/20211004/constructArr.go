package _0211004

func constructArr(a []int) []int {
	if len(a)==0{
		return []int{}
	}
	left:=make([]int,len(a))
	left[0]=1
	for i:=1;i<len(a);i++{
		left[i]=left[i-1]*a[i-1]
	}
	right:=1
	for i:=len(a)-1;i>=0;i--{
		left[i]*=right
		right*=a[i]
	}
	return left
}
