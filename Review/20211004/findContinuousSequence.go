package _0211004

func findContinuousSequence(target int) [][]int {
	res:=[][]int{}
	left,right:=0,0
	for right<target{
		sum:=right*(right+1)/2-left*(left+1)/2
		if sum>target{
			left++
		}else if sum<target{
			right++
		}else{
			temp:=[]int{}
			for i:=left+1;i<=right;i++{
				temp=append(temp,i)
			}
			res=append(res,temp)
			left++
			right++
		}
	}
	return res
}
