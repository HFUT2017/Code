package Baidu

import "strconv"

func BigIntergerAdd(num1 string,num2 string) string{
	l1,l2,cnt:=0,0,0
	res:=""
	for l1<len(num1)||l2<len(num2){
		n1,n2:=0,0
		if l1<len(num1){
			n1=int(num1[l1]-'0')
			l1++
		}
		if l2<len(num2){
			n2=int(num2[l2]-'0')
			l2++
		}
		sum:=(n1+n2+cnt)%10
		cnt=(n1+n2+cnt)/10
		res+=strconv.Itoa(sum)
	}
	if cnt!=0{
		res+=strconv.Itoa(cnt)
	}
	return res
}
