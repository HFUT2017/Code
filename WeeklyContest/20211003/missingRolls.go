package _0211003

func missingRolls(rolls []int, mean int, n int) []int {
	m:=len(rolls)
	sum:=(m+n)*mean
	res:=[]int{}
	for i:=0;i<m;i++{
		sum-=rolls[i]
	}
	num1,num2:=sum/n,sum%n
	if num1>6||num1<=0{
		return []int{}
	}
	for i:=0;i<n;i++{
		if i<num2{
			if num1+1>6{
				return []int{}
			}
			res=append(res,(num1+1))
		}else {
			res=append(res,num1)
		}
	}
	return res
}
