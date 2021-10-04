package _0211004

func singleNumber(nums []int) int {
	one,two:=0,0
	for _,value:=range nums{
		one=one^value&^two
		two=two^value&^one
	}
	return one
}
