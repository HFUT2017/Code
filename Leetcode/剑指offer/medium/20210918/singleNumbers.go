package _0210918

// SingleNumbers 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
func SingleNumbers(nums []int) []int {
	a:=0
	for _,value:=range nums{
		a^=value
	}
	b:=1
	for b&a==0{
		b<<=1
	}
	c,d:=0,0
	for _,value:=range nums{
		if value&b==0{
			c^=value
		}else{
			d^=value
		}
	}
	return []int{c,d}
}
