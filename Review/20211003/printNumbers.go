package _0211003

func printNumbers(n int) []int {
	res:=[]int{}
	ans:=0
	for n!=0{
		ans=ans*10+9
		n--
	}
	for i:=1;i<=ans;i++{
		res=append(res,i)
	}
	return res
}
