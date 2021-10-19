package BDA2

func MergeTwoNUms(num1 []int,num2 []int) []int{
	num3:=make([]int,len(num1)+len(num2))
	l1,l2,tail:=0,0,0
	for l1<len(num1)||l2<len(num2){
		if l1==len(num1){
			num3[tail]=num2[l2]
			l2++
		}else if l2==len(num2){
			num3[tail]=num1[l1]
			l1++
		}else if num1[l1]<num2[l2]{
			num3[tail]=num1[l1]
			l1++
		}else{
			num3[tail]=num2[l2]
			l2++
		}
		tail++
	}
	return num3
}
